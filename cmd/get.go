package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"get"},
	Short:   "package-downloader - a simple CLI to download package from source",
	Long:    `package-downloader - a simple CLI to download package from source`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Name()
		log.Printf("this is get subcommand from Cobra")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
