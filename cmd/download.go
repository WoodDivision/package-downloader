package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Aliases: []string{"download"},
	Short:   "package-downloader - a simple CLI to download package from source",
	Long:    `package-downloader - a simple CLI to download package from source`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("this is download subcommand from Cobra")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
