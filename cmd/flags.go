package cmd

import (
	"fmt"
)

// [name] [flags] [arguments]
type withflags struct {
	basic
	flags []string
}

func (cmd *withflags) AddFlags(flags ...string) {
	cmd.flags = append(cmd.flags, flags...)
}

func (cmd withflags) Run() *Result {
	cmd.arguments = append(cmd.flags, cmd.arguments...)
	return cmd.basic.Run()
}

func (cmd withflags) String() string {
	return fmt.Sprintf("%s %v %v", cmd.name, cmd.flags, cmd.arguments)
}
