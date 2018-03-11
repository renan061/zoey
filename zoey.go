package main

import (
	"fmt"
	"log"

	"github.com/renan061/zoey/cmd"
	"github.com/renan061/zoey/test"
)

func main() {
	// TODO: read from file
	compiler := cmd.New("gcc", 10, "-std=c99", "-Wall")

	assignment, err := test.SetUp(compiler, "./factorial/configuration.json")
	if err != nil {
		log.Fatalln(err)
	}

	var result cmd.Result

	result = assignment.CompileObjects(compiler)
	if !result.Ok {
		fmt.Println("-- Programa não compilou (nota: 0.0)")
		fmt.Println("-- Dump:")
		result.Dump()
		return
	}

	result = assignment.CompileTests(compiler)
	if !result.Ok {
		fmt.Println("-- Erro do teste: testes não compilaram")
		fmt.Println("-- Dump:")
		result.Dump()
		return
	}

	fmt.Println("-- Programa compilou corretamente")

	result = assignment.Run()
	if !result.Ok {
		fmt.Println("-- Erro: não foi possível executar o binário")
		fmt.Println("-- Dump:")
		result.Dump()
		return
	}

	fmt.Println("-- Programa executou todos os testes")

	assignment.TearDown()

	grade, failures := assignment.Grade()
	fmt.Printf("-- Nota final: %d/100\n", grade)
	for _, failure := range failures {
		fmt.Println("- Falhou no teste " + failure)
	}
}
