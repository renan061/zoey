package check

import (
	"fmt"

	"github.com/renan061/zoey/compile"
)

var (
	compiler = "gcc-7"
	timeout  = 10
	stdc     = "-std=c99"
	wall     = "-Wall"
)

type module struct {
	header string
	source string
}

func (m module) compile() (ok bool, stdout, stderr string, err error) {
	cmd := compile.NewCmd(compiler, timeout)
	cmd.AddFlags(stdc, wall, "-iquote"+m.header, "-c")
	cmd.AddArguments(m.source)
	return cmd.Run()
}

// checks the syntax for a single file
func Syntax(file string) (ok bool, message string, err error) {
	compileErrorToString := func(ok bool, stdout, stderr, file string) string {
		if ok {
			return ""
		}
		str := fmt.Sprintf("syntax error in file %s", file)
		if stdout != "" {
			str = fmt.Sprintf("%s\nstdout:\n%s", str, stdout)
		}
		if stderr != "" {
			str = fmt.Sprintf("%s\nstderr:\n%s", str, stderr)
		}
		return str
	}

	cmd := compile.NewCmd("gcc-7", 10)
	cmd.AddFlags("-std=c99", "-Wall", "-fsyntax-only", "-c")
	cmd.AddArguments(file)

	ok, stdout, stderr, err := cmd.Run()
	message = compileErrorToString(ok, stdout, stderr, file)
	return ok, message, err
}

func All(files []string) (ok bool, message string, err error) {
	// checks syntax for all files
	for _, file := range files {
		ok, message, err = Syntax(file)
		if !ok {
			return
		}
	}

	fmt.Println(module{"csrc/stack.h", "csrc/stack.c"}.compile())

	cmd := compile.NewCmd("gcc-7", 10)
	cmd.AddFlags("-std=c99", "-Wall")
	cmd.AddArguments("stack.o", "csrc/main.c")
	fmt.Println(cmd.String())
	fmt.Println(cmd.Run())

	return true, "", nil
}
