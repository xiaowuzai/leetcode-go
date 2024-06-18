package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// [-1,1,2,-2,3]  子串最大和=4
func maxArr(arr []int) int {
	return 0
}
func main() {

	go func() {
		defer func() {
			recover()
		}()
		process()
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	clean()

	fmt.Println(" 程序退出")
}

func process() {}

func clean() {}
