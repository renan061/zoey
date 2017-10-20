package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/renan061/zoey/cmd"
)

var (
	ErrInternal      = errors.New("internal error: test")
	ErrCompilingTest = errors.New("configuration error: could not compile tests")

	ErrCompilingObject = errors.New("compiling error: could not compile object")
)

// TODO
func abs(dir, base string) (string, error) {
	abs, err := filepath.Abs(fmt.Sprintf("%s/%s", dir, base))
	if err != nil {
		fmt.Println("internal error: abs")
		fmt.Println(err)
		return "", ErrInternal
	}
	return abs, nil
}

type (
	// TODO: Doc
	Assignment struct {
		// The name of the assignment.
		Name string `json:"name"`
		// Path to where the assignment directory is in the file system.
		Directory string `json:"directory"`

		Objects []*Object `json:"objects"`

		Tests []*Test `json:"tests"`

		// Auxiliary
		headerdir string
		tempdir   string
		testdir   string
	}

	// TODO: Doc
	Object struct {
		// The name of the object file.
		Name string `json:"name"`
		// The header file containing the exported functions from source.
		Header string `json:"header"`
		// The object implementation file.
		Source string `json:"source"`
	}

	// TODO: Doc
	Test struct {
		// The name of the test (will be visible to the user).
		Name string `json:"name"`
		// A brief description of the test (will be visible to the user).
		Description string `json:"description"`
		// The .c file containing the main function to be executed.
		Main string `json:"main"`
		// The value the main file should print, if any.
		Expected string `json:"expected"`
		// The grade value for the test.
		Value uint `json:"value"`

		// Auxiliary
		binary string
	}
)

func Setup(compiler cmd.Cmd, configuration string) (*Assignment, error) {
	// json unmarshal
	raw, err := ioutil.ReadFile(configuration)
	if err != nil {
		return nil, err
	}
	asgmt := Assignment{}
	err = json.Unmarshal(raw, &asgmt)
	if err != nil {
		return nil, err
	}

	// sets assignment directory to absolute path
	asgmt.Directory, err = filepath.Abs(asgmt.Directory)
	if err != nil {
		fmt.Println(err)
		return nil, ErrInternal
	}

	// sets auxiliary directory names
	asgmt.headerdir = asgmt.Directory + "/headers"
	asgmt.tempdir = asgmt.Directory + "/temp"
	asgmt.testdir = asgmt.Directory + "/tests"

	// adds include search path flag (headerdir)
	compiler.AddFlags("-I" + asgmt.headerdir)

	// creates a temporary directory
	err = createDirectory(asgmt.tempdir)
	if err != nil {
		return nil, err
	}

	// sets header and source paths to absolute path
	for _, object := range asgmt.Objects {
		object.Name = asgmt.tempdir + "/" + object.Name
		object.Header = asgmt.Directory + "/" + object.Header
		object.Source = asgmt.Directory + "/" + object.Source
	}

	// sets main paths to absolute path
	for _, test := range asgmt.Tests {
		test.Main = asgmt.testdir + "/" + test.Main
		test.binary = asgmt.tempdir + "/" + test.Name
	}

	return &asgmt, nil
}

func (asgmt *Assignment) CompileObjects(compiler cmd.Cmd) (*cmd.Result, error) {
	// TODO: Run in parallel
	// compiles the object files
	for _, object := range asgmt.Objects {
		result, err := compileObject(compiler, object.Source, object.Name)
		if err != nil {
			return result, err
		}
	}

	return nil, nil
}

func (asgmt *Assignment) CompileTests(compiler cmd.Cmd) (*cmd.Result, error) {
	// TODO: Create function
	// ex.: gcc temp/factorial main.c -o temp/assert.out
	objects := []string{}
	for _, object := range asgmt.Objects {
		objects = append(objects, object.Name)
	}

	for _, test := range asgmt.Tests {
		compiler.AddArguments(objects...)
		compiler.AddArguments(test.Main, "-o", test.binary)
		result := compiler.Run()
		if !result.Ok {
			result.Dump()
			return result, ErrCompilingTest
		}
		compiler.ResetArguments()
	}

	return nil, nil
}

// func (t Test) Run() error {
// 	file, err := abs(t.Directory, "./temp/"+assertout)
// 	if err != nil {
// 		return err
// 	}

// 	// ./temp/assert.out
// 	result := cmd.New(file, 10).Run()
// 	result.Dump() // TODO: ?
// 	return nil
// }

func (asgmt Assignment) Teardown() error {
	return removeDirectory(asgmt.tempdir)
}

// ==================================================
//
//	Auxiliary
//
// ==================================================

// - path is the absolute path for the new directory
func createDirectory(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Println("internal error: os.MkdirAll") // TODO
		return ErrInternal
	}

	return nil
}

// removes a directory (deleting all that is inside it)
// - path is the absolute path to the directory that will be removed
func removeDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("internal error: os.RemoveAll")
		return ErrInternal
	}

	return nil
}

// ex.: gcc -c src.c -o obj
// - src is the absolute path to the source code to be compiled
// - obj is the absolute path to where the object should be created
func compileObject(compiler cmd.Cmd, src, obj string) (*cmd.Result, error) {
	compiler.AddArguments("-c", src, "-o", obj)
	result := compiler.Run()
	if !result.Ok {
		result.Dump()
		return result, ErrCompilingObject
	}
	compiler.ResetArguments()

	return result, nil
}

// returns the name of an object file given its source file
func objname(src string) string {
	return strings.Split(filepath.Base(src), ".")[0]
}
