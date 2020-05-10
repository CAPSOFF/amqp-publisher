package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(os.Geteuid())
	fmt.Println("test 123")
}
