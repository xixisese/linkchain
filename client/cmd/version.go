package cmd

import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  "This is get the version number of linkchain command",
	Run: func(cmd *cobra.Command, args []string) {
		println("linkchain version is 0.0.1")
	},
}
