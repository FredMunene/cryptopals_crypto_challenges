package main

import "strconv"

func main(){
	println(toBinary(73))
}

func toBinary(number int) string{
	return strconv.FormatInt(int64(number),2)
}