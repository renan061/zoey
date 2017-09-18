package main

import (
	"fmt" // TODO: Remove

	"github.com/renan061/zoey/check"
)

func main() {
	files := []string{"csrc/stack.h", "csrc/stack.c", "csrc/main.c"}
	ok, message, err := check.All(files)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !ok {
		fmt.Printf("Message: %v\n", message)
		return
	}
}
