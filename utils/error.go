package utils


func DoError (err error)  {
	if(err != nil){
		panic(err)
	}
}