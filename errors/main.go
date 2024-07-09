package main

import "fmt"

type error interface {
	Error() string
}

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg + ": Error"
}

func Log(e error) {
	fmt.Print(e)
}

func main() {
	e := MyError{"Hello world"}
	Log(e)
}
