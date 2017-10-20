package cmd

import (
	"fmt"
)

type (
	Cmd interface {
		Run() *Result
		AddArguments(...string)
		AddFlags(...string)
		ResetArguments()
		String() string
	}

	Result struct {
		Ok     bool
		StdOut StdOut
		StdErr StdErr
		Err    error
	}

	StdOut string

	StdErr string
)

func New(name string, timeout int, flags ...string) Cmd {
	if len(flags) < 1 {
		return &basic{name: name, timeout: timeout}
	} else {
		return &withflags{
			basic: basic{name: name, timeout: timeout},
			flags: flags,
		}
	}
}

func (result Result) Dump() {
	if result.Err != nil {
		fmt.Println("internal error: cmd")
		fmt.Println(result.Err)
		return
	}

	if !result.Ok {
		fmt.Println("test failed")
	}

	result.StdOut.Dump()
	result.StdErr.Dump()
}

func (stdout StdOut) Dump() {
	str := string(stdout)
	if str != "" {
		fmt.Printf("stdout: \n%s\n", str)
	} else {
		fmt.Printf("stdout is empty\n")
	}
}

func (stderr StdErr) Dump() {
	str := string(stderr)
	if str != "" {
		fmt.Printf("stderr: %s\n", str)
	} else {
		fmt.Printf("stderr is empty\n")
	}
}
