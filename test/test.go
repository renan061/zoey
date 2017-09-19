package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/renan061/zoey/cmd"
)

type (
	Test struct {
		Name      string `json:"name"`
		Directory string `json:"directory"`
		Assert    Assert `json:"assert"`
	}

	Assert struct {
		Main    string   `json:"main"`
		Objects []Object `json:"objects"`
	}

	Object struct {
		Header string `json:"header"`
		Source string `json:"source"`
	}
)

func LoadFrom(file string) (*Test, error) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	test := Test{}
	err = json.Unmarshal(raw, &test)
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func (t Test) Run() {
	const timeout = 10

	gcc := cmd.New("gcc-7", timeout, "-std=c99", "-Wall")

	pathto := func(file string) string {
		return fmt.Sprintf("%s/%s", t.Directory, file)
	}

	objname := func(i int) string {
		return fmt.Sprintf("obj%d", i)
	}

	check := func(result *cmd.Result) bool {
		if result.Err != nil {
			fmt.Println(result.Err)
			return false
		}

		if !result.Ok {
			fmt.Printf("test failed\n")
			if result.StdOut != "" {
				fmt.Printf("stdout: %s\n", result.StdOut)
			}
			if result.StdErr != "" {
				fmt.Printf("stderr: %s\n", result.StdErr)
			}
			return false
		}

		if result.StdOut != "" {
			fmt.Printf("stdout:\n%s", result.StdOut)
		}
		if result.StdErr != "" {
			fmt.Printf("stderr:\n%s", result.StdErr)
		}
		return true
	}

	// gcc -c ex/object.c -o ex/object.o
	for i, object := range t.Assert.Objects {
		gcc.AddArguments("-c", pathto(object.Source), "-o", pathto(objname(i)))
		result := gcc.Run()
		if !check(result) {
			return
		}
		gcc.ResetArguments()
	}

	// gcc ex/obj0.o ex/assert.c -o ex/assert.out
	for i := 0; i < len(t.Assert.Objects); i++ {
		gcc.AddArguments(pathto(objname(i)))
	}
	const assertout = "assert.out"
	gcc.AddArguments(pathto(t.Assert.Main), "-o", pathto(assertout))
	result := gcc.Run()
	if !check(result) {
		return
	}
	gcc.ResetArguments()

	// ./ex/assert.out
	run := cmd.New(pathto(assertout), timeout)
	result = run.Run()
	if !check(result) {
		return
	}
}
