package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)

var openJiraCommand = &cobra.Command{
	Use:   "open-jira",
	Short: "open jira ticket",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		envs, loadError := godotenv.Read()

		if loadError != nil {
			log.Fatal("Can not load any environment variables")
		}

		var jiraSite = args[0]
		// Need to replace jira link
		var shell = exec.Command("open", fmt.Sprintf("%s%s", envs["JIRA_SITE"], jiraSite))
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
