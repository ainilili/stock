package logger

import "fmt"

func Error(msg interface{}) {
	fmt.Println(msg)
}

func Errorf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}
