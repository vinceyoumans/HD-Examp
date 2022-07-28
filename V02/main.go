package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("start here homedepot V02")
	fmt.Println("Hello there!")

	fmt.Println("multi-stage build")

	osv := runtime.GOOS
	fmt.Printf("The OS is: %s\n", osv)
}
