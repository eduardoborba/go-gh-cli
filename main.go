package main

import (
	"github.com/eduardoborba/go-gh-cli/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "issues"}
	rootCmd.AddCommand(cmd.ListIssuesCmd)
	rootCmd.Execute()
}
