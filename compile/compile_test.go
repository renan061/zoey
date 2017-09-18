package compile

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompileCmd(t *testing.T) {
	require := require.New(t)

	const (
		compiler = "gcc"
		timeout  = 5
		flag     = "-flag"
		argument = "argument"
	)

	// NewCmd
	cmd := NewCmd(compiler, timeout)
	require.NotNil(cmd)
	require.Equal(cmd.compiler, compiler)
	require.Equal(cmd.timeout, timeout)

	// AddFlags
	for i := 0; i < 10; i++ {
		cmd.AddFlags(flag + string(i))
	}
	for i := 0; i < 10; i++ {
		require.Equal(cmd.flags[i], flag+string(i))
	}

	// AddArguments
	for i := 0; i < 10; i++ {
		cmd.AddArguments(argument + string(i))
	}
	for i := 0; i < 10; i++ {
		require.Equal(cmd.arguments[i], argument+string(i))
	}

	// Run - Ok
	cmd = NewCmd(compiler, timeout)
	cmd.AddFlags("-std=c99", "-Wall", "-o")
	cmd.AddArguments("/dev/null", "../resources/compile_test_ok.c")
	ok, stdout, stderr, err := cmd.Run()
	require.Nil(err)
	require.True(ok)
	require.Empty(stdout)
	require.Empty(stderr)

	// Run - Error
	cmd = NewCmd(compiler, timeout)
	cmd.AddFlags("-std=c99", "-Wall", "-o")
	cmd.AddArguments("/dev/null", "../resources/compile_test_error.c")
	ok, stdout, stderr, err = cmd.Run()
	require.Nil(err)
	require.False(ok)
	require.Empty(stdout)
	require.NotEmpty(stderr)
}
