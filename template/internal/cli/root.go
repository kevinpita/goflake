// Package cli owns command-line parsing and application output.
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"example.com/project/internal/app"
)

// NewCommand builds the application's command tree.
func NewCommand() *cobra.Command {
	var name string
	root := &cobra.Command{
		Use:           "app",
		Short:         "A Go application starter",
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), app.Greeting(name))
			return err
		},
	}
	root.Flags().StringVar(&name, "name", "world", "Name to greet")
	return root
}
