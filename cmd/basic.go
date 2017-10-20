package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

// [name] [arguments]
type basic struct {
	name      string
	timeout   int
	arguments []string
}

func (cmd *basic) AddArguments(arguments ...string) {
	cmd.arguments = append(cmd.arguments, arguments...)
}

func (cmd *basic) AddFlags(flags ...string) {
}

func (cmd *basic) ResetArguments() {
	cmd.arguments = []string{}
}

func (cmd basic) Run() *Result {
	timeout := time.Duration(cmd.timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	execCmd := exec.CommandContext(ctx, cmd.name, cmd.arguments...)

	result := &Result{}

	stderrPipe, err := execCmd.StderrPipe()
	if result.Err = err; err != nil {
		return result
	}
	stdoutPipe, err := execCmd.StdoutPipe()
	if result.Err = err; err != nil {
		return result
	}
	execCmd.Start()

	stderrBytes, err := ioutil.ReadAll(stderrPipe)
	if result.Err = err; err != nil {
		return result
	}
	stdoutBytes, err := ioutil.ReadAll(stdoutPipe)
	if result.Err = err; err != nil {
		return result
	}
	execCmd.Wait()

	result.StdOut = StdOut(stdoutBytes)
	result.StdErr = StdErr(stderrBytes)
	result.Ok = execCmd.ProcessState.Success() && result.StdErr == ""
	result.Err = nil

	return result
}

func (cmd basic) String() string {
	return fmt.Sprintf("%s %v", cmd.name, cmd.arguments)
}
