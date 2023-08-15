package cmd

import (
	"github.com/spf13/cobra"
	"os/exec"
)

var openJiraCommand = &cobra.Command{
	Use:   "open-jira",
	Short: "open jira ticket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var jiraSite = args[0]
		// Need to replace jira link
		var shell = exec.Command("open", "https://www.naver.com")
		err := shell.Run()
		if err != nil {
			return
		}
		println(jiraSite)
	},
}

func init() {
	rootCmd.AddCommand(openJiraCommand)
}
