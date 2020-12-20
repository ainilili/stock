/**
2 * @Author: Nico
3 * @Date: 2020/12/21 0:45
4 */
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
