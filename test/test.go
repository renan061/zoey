package test

// TODO: run compilations in parallel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/renan061/zoey/cmd"
)

var (
	ErrInternal = errors.New("internal error: test")
)

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
		// The object implementation file.
		Source string `json:"source"`
	}

	// TODO: Doc
	Test struct {
		// The name of the test (will be visible to the user).
		Name string `json:"name"`
		// A brief description of the test (will be visible to the user).
		Description string `json:"description"`
		// The .c file containing the "main" function.
		Main string `json:"main"`
		// The value the main file should print, if any.
		Expected string `json:"expected"`
		// The grade value for the test.
		Value uint `json:"value"`

		// Auxiliary
		binary string
		output string
		ok     bool
	}
)

func SetUp(compiler cmd.Cmd, configuration string) (*Assignment, error) {
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
		object.Source = asgmt.Directory + "/" + object.Source
	}

	// sets main paths to absolute path
	for _, test := range asgmt.Tests {
		test.Main = asgmt.testdir + "/" + test.Main
		test.binary = asgmt.tempdir + "/" + test.Name + ".out"
	}

	return &asgmt, nil
}

func (asgmt Assignment) CompileObjects(compiler cmd.Cmd) cmd.Result {
	// ex.: gcc -c src.c -o obj
	// - c is the command for the compiler
	// - src is the absolute path to the source code to be compiled
	// - obj is the absolute path to where the object should be created
	compile := func(c cmd.Cmd, src, obj string) cmd.Result {
		c.AddArguments("-c", src, "-o", obj)
		result := c.Run()
		if result.Ok {
			c.ResetArguments()
		}
		return result
	}

	// compiles the objects
	for _, object := range asgmt.Objects {
		result := compile(compiler, object.Source, object.Name)
		if !result.Ok {
			return result
		}
	}

	return cmd.Result{Ok: true}
}

func (asgmt Assignment) CompileTests(compiler cmd.Cmd) cmd.Result {
	// ex.: gcc [obj1 obj2 ...] main.c -o bin.out
	// - c is the command for the compiler
	// - objs is a list with absolute paths to where the objects are located
	// - main is the absolute path to the source with the "main" function
	// - bin is the absolute path to where the binary should be created
	compile := func(c cmd.Cmd, objs []string, main, bin string) cmd.Result {
		c.AddArguments(objs...)
		c.AddArguments(main, "-o", bin)
		result := c.Run()
		if result.Ok {
			c.ResetArguments()
		}
		return result
	}

	// creates a list with all the objects' absolute paths
	objects := []string{}
	for _, object := range asgmt.Objects {
		objects = append(objects, object.Name)
	}

	// compiles the tests
	for _, test := range asgmt.Tests {
		result := compile(compiler, objects, test.Main, test.binary)
		if !result.Ok {
			return result
		}
	}

	return cmd.Result{Ok: true}
}

func (asgmt Assignment) Run() cmd.Result {
	// runs the tests
	for _, test := range asgmt.Tests {
		result := cmd.New(test.binary, 10).Run()
		if !result.Ok {
			return result
		}
		test.output = string(result.StdOut)
		test.ok = test.output == test.Expected
	}

	return cmd.Result{Ok: true}
}

func (asgmt Assignment) TearDown() error {
	return removeDirectory(asgmt.tempdir)
}

func (asgmt Assignment) Grade() (uint, []string) {
	failures := []string{}
	total := uint(0)
	grade := uint(0)
	for _, test := range asgmt.Tests {
		total += test.Value
		if test.ok {
			grade += test.Value
		} else {
			format := "'%s' (esperado: '%s', obtido: '%s')"
			str := fmt.Sprintf(format, test.Name, test.Expected, test.output)
			failures = append(failures, str)
		}
	}
	grade = uint(float32(grade) / float32(total) * 100)
	return grade, failures
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
