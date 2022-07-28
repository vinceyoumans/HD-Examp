package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("start here   homedepot V01")
	fmt.Println("Hello there!")

	osv := runtime.GOOS
	fmt.Printf("The OS is: %s\n", osv)

}
