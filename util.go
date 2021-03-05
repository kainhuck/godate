package godate

func assert(condition bool, msg string) {
	if !condition {
		panic(msg)
	}
}


func wrapFuncIntErr(i int, e error) int{
	if e != nil{
		panic(e)
	}

	return i
}