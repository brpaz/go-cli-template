package main

import (
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/brpaz/go-cli-template/cmd/hello"
	versionCmd "github.com/brpaz/go-cli-template/cmd/version"
)

var (
	version   = "dev"
	gitCommit = "none"
	buildDate = "unknown"
)

// NewRootCommand creates a new instance of the root command
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "hello",
		Short:   "An example CLI application",
		Version: version,
	}

	return rootCmd
}

func registerCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(hello.NewCommand())
	rootCmd.AddCommand(versionCmd.NewCommand(versionCmd.VersionInfo{
		Version:   version,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
	}))
}

// Application entrypoint
func main() {
	rootCmd := NewRootCommand()
	registerCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
