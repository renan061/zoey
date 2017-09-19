package main

import (
	"log"

	"github.com/renan061/zoey/test"
)

func main() {
	// files := []string{"ex/stack.h", "ex/stack.c", "ex/main.c"}
	// ok, message, err := check.All(files)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// if !ok {
	// 	fmt.Printf("Message: %v\n", message)
	// 	return
	// }

	t, err := test.LoadFrom("ex/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	t.Run()
}
