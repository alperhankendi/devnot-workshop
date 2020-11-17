package cmd

import (
	"fmt"
	"github.com/alperhankendi/devnot-workshop/internal/movies"
	"github.com/alperhankendi/devnot-workshop/internal/movies/controller"
	"github.com/alperhankendi/devnot-workshop/internal/movies/mongo"
	"github.com/alperhankendi/devnot-workshop/pkg/echoextention"
	"github.com/alperhankendi/devnot-workshop/pkg/log"
	mongohelper "github.com/alperhankendi/devnot-workshop/pkg/mongoextentions"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"time"
)

type queryCommand struct {
	instance *echo.Echo
	command  *cobra.Command
}

var queryCmd = &queryCommand{
	command: &cobra.Command{
		Use:   "query",
		Short: "Query Service",
	},
}

func init() {
	rootCmd.AddCommand(queryCmd.command)

	queryCmd.instance = echo.New()
	queryCmd.instance.Debug = false
	queryCmd.instance.HidePort = true
	queryCmd.instance.HideBanner = true
	queryCmd.instance.Logger = log.SetupLogger()
	echoextention.RegisterGlobalMiddlewares(queryCmd.instance)
	queryCmd.command.RunE = func(cmd *cobra.Command, args []string) error {

		var repo movies.Repository
		mongoConn, err := mongohelper.NewDatabase(dbConn, dbName)
		if err != nil {
			log.Logger.Fatalf("Failed to connect database. Error:%v", err)
		}
		repo = mongo.NewRepository(mongoConn)
		service := movies.NewService(repo)
		rs := controller.NewController(service)
		controller.BuildQueryHandler(queryCmd.instance, rs)
		log.Logger.Infof("Service is starting. Service port:%s", port)
		go func() {
			if err := queryCmd.instance.Start(fmt.Sprintf(":%s", port)); err != nil {
				log.Logger.Fatalf("Failed to shutting down the server. Error :%v", err)
			}
		}()
		echoextention.Shutdown(queryCmd.instance, time.Second*3)

		return nil
	}
}
