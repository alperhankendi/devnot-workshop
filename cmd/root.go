package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "devnot-workshop",
	Short: "Devnote workshop application",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	port   string
	dbConn string
	dbName string
)

func init() {

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "5001", "Service Port")
	rootCmd.PersistentFlags().StringVarP(&dbConn, "conn", "c", "mongodb://root:example@localhost:27017", "database connection string")
	rootCmd.PersistentFlags().StringVarP(&dbName, "dbname", "d", "imdb", "database name")
}
