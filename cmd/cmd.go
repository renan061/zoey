package cmd

import (
	"fmt"
)

type (
	Cmd interface {
		Run() Result
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
	basic := &basic{name: name, timeout: timeout}
	if len(flags) < 1 {
		return basic
	} else {
		return &withflags{basic: *basic, flags: flags}
	}
}

func (result Result) Dump() {
	if result.Err != nil {
		fmt.Println("internal error: cmd")
		fmt.Println(result.Err)
		return
	}

	if !result.Ok {
		fmt.Println("cmd.Ok: false")
	}

	result.StdOut.Dump()
	result.StdErr.Dump()
}

func (stdout StdOut) Dump() {
	str := string(stdout)
	if str != "" {
		fmt.Printf("cmd.StdOut: \n%s\n", str)
	} else {
		fmt.Printf("cmd.StdOut: empty\n")
	}
}

func (stderr StdErr) Dump() {
	str := string(stderr)
	if str != "" {
		fmt.Printf("cmd.StdErr: %s\n", str)
	} else {
		fmt.Printf("cmd.StdErr: empty\n")
	}
}
