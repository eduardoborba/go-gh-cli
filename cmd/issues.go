package cmd

import (
	"fmt"

	"github.com/eduardoborba/go-gh-cli/logic"
	"github.com/spf13/cobra"
)

var ListIssuesCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetches the issues for a github repo",
	Long:  `Feches and prints the info issues for a given github repo`,
	Run: func(cmd *cobra.Command, args []string) {
		issues, err := logic.FetchIssues("https://api.github.com")
		if err != nil {
			panic(err)
		}

		for _, issue := range issues {
			fmt.Println("id:\t", issue.ID)
			fmt.Println("url:\t", issue.Url)
		}
	},
}
