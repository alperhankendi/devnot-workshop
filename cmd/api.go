package cmd

import (
	"github.com/spf13/cobra"
)

var (
	port   string
	dbConn string
	dbName string
)
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "select api type",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.PersistentFlags().StringVarP(&port, "port", "p", "5001", "Service Port")
	apiCmd.PersistentFlags().StringVarP(&dbConn, "conn", "c", "mongodb://root:example@localhost:27017", "database connection string")
	apiCmd.PersistentFlags().StringVarP(&dbName, "dbname", "d", "imdb", "database name")
}
