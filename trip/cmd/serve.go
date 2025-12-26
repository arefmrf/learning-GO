package cmd

import (
	"github.com/spf13/cobra"
	"trip/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve backend",
	Long:  "serve backend",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Run()
}
