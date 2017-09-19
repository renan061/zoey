package cmd

type (
	Cmd interface {
		Run() *Result
		AddArguments(...string)
		ResetArguments()
		String() string
	}

	Result struct {
		Ok     bool
		StdOut string
		StdErr string
		Err    error
	}
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
