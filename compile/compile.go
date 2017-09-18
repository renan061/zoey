package compile

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

type Cmd struct {
	compiler  string
	timeout   int
	flags     []string
	arguments []string
}

func NewCmd(compiler string, timeout int) *Cmd {
	return &Cmd{compiler: compiler, timeout: timeout}
}

func (cmd *Cmd) AddFlags(flags ...string) {
	cmd.flags = append(cmd.flags, flags...)
}

func (cmd *Cmd) AddArguments(arguments ...string) {
	cmd.arguments = append(cmd.arguments, arguments...)
}

func (cmd Cmd) Run() (ok bool, stdout, stderr string, err error) {
	timeout := time.Duration(cmd.timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	args := append(cmd.flags, cmd.arguments...)
	execCmd := exec.CommandContext(ctx, cmd.compiler, args...)

	stderrPipe, err := execCmd.StderrPipe()
	if err != nil {
		return false, "", "", err
	}
	stdoutPipe, err := execCmd.StdoutPipe()
	if err != nil {
		return false, "", "", err
	}
	execCmd.Start()

	stderrBytes, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		return false, "", "", err
	}
	stdoutBytes, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return false, "", "", err
	}
	execCmd.Wait()

	ok = execCmd.ProcessState.Success()
	stdout = string(stdoutBytes)
	stderr = string(stderrBytes)
	return ok, stdout, stderr, nil
}

func (cmd Cmd) String() string {
	return fmt.Sprintf("%s %v %v", cmd.compiler, cmd.flags, cmd.arguments)
}
