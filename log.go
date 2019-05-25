package mvc

import (
	"log"
	"os"
)

const(
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var level = LevelTrace

func Level() int {
	return level
}

func SetLevel(l int) {
	level = l
}

var ChanuLogger  = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func SetLogger(l *log.Logger){
	ChanuLogger = l
}

func Trace(v ...interface{}){
	if level <= LevelTrace {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}
func Debug(v ...interface{}){
	if level <= LevelDebug {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}
func Info(v ...interface{}){
	if level <= LevelInfo {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}
func Warning(v ...interface{}){
	if level <= LevelWarning {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}
func Error(v ...interface{}){
	if level <= LevelError {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}
func Critical(v ...interface{}){
	if level <= LevelCritical {
		ChanuLogger.Printf("[T] %v¥n",v)
	}
}



