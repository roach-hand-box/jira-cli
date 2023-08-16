package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var openJiraCommand = &cobra.Command{
	Use:   "open-jira",
	Short: "open jira ticket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var jiraSite = args[0]
		// Need to replace jira link
		var shell = exec.Command("open", fmt.Sprintf("%s%s", os.Getenv("JIRA_SITE"), jiraSite))
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
