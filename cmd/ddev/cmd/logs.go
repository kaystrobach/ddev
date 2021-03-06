package cmd

import (
	"strings"

	"github.com/drud/ddev/pkg/ddevapp"
	"github.com/drud/ddev/pkg/util"
	"github.com/spf13/cobra"
)

var (
	tail      string
	follow    bool
	timestamp bool
)

// DdevLogsCmd contains the "ddev logs" command
var DdevLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get the logs from your running services.",
	Long:  `Uses 'docker logs' to display stdout from the running services.`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := ddevapp.GetActiveApp("")
		if err != nil {
			util.Failed("Failed to retrieve logs: %v", err)
		}

		if strings.Contains(app.SiteStatus(), ddevapp.SiteNotFound) {
			util.Failed("Project is not currently running. Try 'ddev start'.")
		}

		if strings.Contains(app.SiteStatus(), ddevapp.SiteStopped) {
			util.Failed("Project is stopped. Run 'ddev start' to start the environment.")
		}

		err = app.Logs(serviceType, follow, timestamp, tail)
		if err != nil {
			util.Failed("Failed to retrieve logs for %s: %v", app.GetName(), err)
		}
	},
}

func init() {
	DdevLogsCmd.Flags().BoolVarP(&follow, "follow", "f", false, "Follow the logs in real time.")
	DdevLogsCmd.Flags().BoolVarP(&timestamp, "time", "t", false, "Add timestamps to logs")
	DdevLogsCmd.Flags().StringVarP(&serviceType, "service", "s", "web", "Defines the service to retrieve logs from. [e.g. web, db]")
	DdevLogsCmd.Flags().StringVarP(&tail, "tail", "", "", "How many lines to show")
	RootCmd.AddCommand(DdevLogsCmd)

}
