package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hello <something>")
		return
	}
	fmt.Printf("Hello %s\n", os.Args[1])
}
