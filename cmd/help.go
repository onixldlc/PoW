package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func (c *CLI) customUsage() {

	// get filename of the executable without the path in both windows and unix
	executable := os.Args[0]
	if os.PathSeparator == '\\' {
		executable = executable[strings.LastIndex(executable, "\\")+1:]
	} else {
		executable = executable[strings.LastIndex(executable, "/")+1:]
	}

	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", executable)
	fmt.Fprintf(flag.CommandLine.Output(), " %s <Commands> [host|-p] [-p <port>]\n", executable)
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Commands:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  s, sender       Use sender mode\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  r, receiver     Use receiver mode\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  m, middleware   Use middleware mode\n")
	fmt.Fprintf(flag.CommandLine.Output(), "\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  [host]\n")
	fmt.Fprintf(flag.CommandLine.Output(), "      Bind address (default 0.0.0.0)\n")
	fmt.Fprintf(flag.CommandLine.Output(), "  -p <port>\n")
	fmt.Fprintf(flag.CommandLine.Output(), "      Bind port (default 8080)\n\n")
}
