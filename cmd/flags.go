package cmd

import (
	"fmt"
)

type withflags struct {
	basic
	flags []string
}

func (cmd withflags) Run() *Result {
	cmd.arguments = append(cmd.flags, cmd.arguments...)
	return cmd.basic.Run()
}

func (cmd withflags) String() string {
	return fmt.Sprintf("%s %v %v", cmd.name, cmd.flags, cmd.arguments)
}
