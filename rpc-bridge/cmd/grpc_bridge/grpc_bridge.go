package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello world '24!")
		time.Sleep(time.Second * 2)
	}
}
