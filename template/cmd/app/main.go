// Command app is the starter application's command-line entry point.
package main

import (
	"fmt"
	"os"

	"example.com/project/internal/cli"
)

func main() {
	if err := cli.NewCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "app:", err)
		os.Exit(1)
	}
}
