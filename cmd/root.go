// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"os"

	"github.com/devops-kung-fu/hinge/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.1"
	verbose bool
	rootCmd = &cobra.Command{
		Use:     "hinge [flags] path/to/repo",
		Example: "  hinge path/to/repo",
		Short:   "Creates and updates your Dependabot config.",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				color.Style{color.FgRed, color.OpBold}.Println("Please provide the path to the repo.")
				fmt.Println()
				cmd.Usage()
			} else if len(args) > 1 {
				color.Style{color.FgRed, color.OpBold}.Println("Only one path is allowed.")
				fmt.Println()
				cmd.Usage()
			} else {
				repoPath := args[0]
				lib.Generator(repoPath)
			}
		},
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	color.Style{color.FgWhite, color.OpBold}.Println("Hinge")
	fmt.Println("https://github.com/devops-kung-fu/hinge")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("")
}
