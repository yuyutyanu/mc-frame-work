package mc_frame_work

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

var (
	cComment = []byte{'#'}
	cEmpty   = []byte{}
	cEqual   = []byte{'='}
	cDQuote  = []byte{'"'}
)

type Config struct {
	filename string
	comment  map[int][]string
	data     map[string]string
	offset   map[string]int64
	mutex    sync.RWMutex
}

func LoadConfig(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{
		filename: file.Name(),
		comment:  make(map[int][]string),
		data:     make(map[string]string),
		offset:   make(map[string]int64),
		mutex:    sync.RWMutex{},
	}
	cfg.mutex.Lock()
	defer cfg.mutex.Unlock()

	var comment bytes.Buffer
	buf := bufio.NewReader(file)

	for nComment, off := 0, int64(1); ; {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		if bytes.Equal(line, cEmpty) {
			continue
		}

		off += int64(len(line))
		// comment行の選別
		if bytes.HasPrefix(line, cComment) {
			line = bytes.TrimLeft(line, "#")
			line = bytes.TrimLeftFunc(line, unicode.IsSpace)
			comment.Write(line)
			comment.WriteByte('\n')
			continue
		}
		if comment.Len() != 0 {
			cfg.comment[nComment] = []string{comment.String()}
			comment.Reset()
			nComment++
		}

		val := bytes.SplitN(line, cEqual, 2)
		if bytes.HasPrefix(val[1], cDQuote) {
			val[1] = bytes.Trim(val[1], `"`)
		}

		key := strings.TrimSpace(string(val[0]))
		cfg.comment[nComment-1] = append(cfg.comment[nComment-1], key)
		cfg.data[key] = strings.TrimSpace(string(val[1]))
		cfg.offset[key] = off
	}

	return cfg, nil
}

func (c *Config) Bool(key string) (bool, error) {
	return strconv.ParseBool(c.data[key])
}

func (c *Config) Int(key string) (int, error) {
	return strconv.Atoi(c.data[key])
}

func (c *Config) Float(key string) (float64, error) {
	return strconv.ParseFloat(c.data[key], 64)
}

func (c *Config) String(key string) string {
	return c.data[key]
}

func (c *Config) GetComment() map[int][]string{
	return c.comment
}