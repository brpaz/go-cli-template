package hello

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCommand returns a new cobra command that prints hello world
func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "hello",
		Short: "Print hello world",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}

	// TODO add flags here

	return command
}

// run executes the sync command
func run(cmd *cobra.Command) error {
	out := cmd.OutOrStdout()
	fmt.Fprintf(out, "Hello World\n")

	return nil
}
