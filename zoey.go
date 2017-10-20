package main

import (
	"fmt"
	"log"

	"github.com/renan061/zoey/cmd"
	"github.com/renan061/zoey/test"
)

func main() {
	// TODO: Constants or read from file
	compiler := cmd.New("gcc-7", 10, "-std=c99", "-Wall")

	assignment, err := test.Setup(compiler, "./factorial/configuration.json")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = assignment.CompileObjects(compiler)
	if err != nil {
		fmt.Println(err)
		fmt.Println("-- Programa não compilou (nota: 0.0)")
		return
	}

	_, err = assignment.CompileTests(compiler)
	if err != nil {
		fmt.Println(err)
		fmt.Println("-- Erro interno: testes não compilaram")
		return
	}

	fmt.Println("-- Programa compilou corretamente")
	// t.Run()
	assignment.Teardown()
}
