package v2

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	v2 "toscactl/tosca/v2"
)

var projectCmd = &cobra.Command{
	Use:   `project`,
	Short: "Show Projects info",
	Long:  `Show Projects info`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := v2.GetProjects()
		project, err := io.ReadAll(resp.Body)
		if err == nil {
			fmt.Printf("Projects: %s Status: %s\"", project, resp.Status)
		}
		os.Exit(1)
	},
}

func init() {
	RootCmd.AddCommand(projectCmd)
}
