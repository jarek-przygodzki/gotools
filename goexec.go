// Execute command with timeout

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

var GitCommit string

func main() {

	timeoutPtr := flag.Duration("timeout", 1*time.Minute, "Command timeout")
	commandPtr := flag.String("command", "", "Command")
	verbosePtr := flag.Bool("verbose", false, "Verbose output")
	versionPtr := flag.Bool("version", false, "Version info")

	flag.Parse()

	timeout := *timeoutPtr
	command := *commandPtr
	verbose := *verbosePtr
	version := *versionPtr

	if version {
		fmt.Printf("goexec version: %s\n", GitCommit)
		return
	}

	if command == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	args := flag.Args()

	if verbose {
		fmt.Printf("Executing command '%s' with args '%s' and timeout %s\n", command, args, timeout)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmd := exec.CommandContext(ctx, command, args...)
	// Pipe stdin and stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Fprintf(os.Stderr, "Failed to execute command with timeout. Command timed out.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command with timeout. Non-zero exit code: %v.\n", err)
		if exiterr, ok := err.(*exec.ExitError); ok {
			// The program has exited with an exit code != 0
			// This works on both Unix and Windows. Although package
			// syscall is generally platform dependent, WaitStatus is
			// defined for both Unix and Windows and in both cases has
			// an ExitStatus() method with the same signature.
			// See https://stackoverflow.com/questions/10385551/get-exit-code-go
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				if verbose {
					fmt.Fprintf(os.Stderr, "Exit code: %d\n", status.ExitStatus())
				}
				os.Exit(status.ExitStatus())
			}
		}
		os.Exit(1)
	}
}
