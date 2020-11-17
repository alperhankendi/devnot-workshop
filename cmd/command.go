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
	"time"

	"github.com/spf13/cobra"
)

type cmdCommand struct {
	instance *echo.Echo
	command  *cobra.Command
}

var cmd = &cmdCommand{
	command: &cobra.Command{
		Use:   "command",
		Short: "Command Service",
	},
}

func init() {
	rootCmd.AddCommand(cmd.command)

	cmd.instance = echo.New()
	cmd.instance.Logger = log.SetupLogger()
	echoextention.RegisterGlobalMiddlewares(cmd.instance)
	cmd.command.RunE = func(c *cobra.Command, args []string) error {

		var repo movies.Repository
		mongoConn, err := mongohelper.NewDatabase(dbConn, dbName)
		if err != nil {
			log.Logger.Fatalf("Failed to connect database. Error:%v", err)
		}
		repo = mongo.NewRepository(mongoConn)
		service := movies.NewService(repo)
		rs := controller.NewController(service)
		controller.BuildCommandHandler(cmd.instance, rs)
		log.Logger.Infof("Service is starting. Service port:%s", port)
		go func() {
			if err := cmd.instance.Start(fmt.Sprintf(":%s", port)); err != nil {
				log.Logger.Fatalf("Failed to shutting down the server. Error :%v", err)
			}
		}()
		echoextention.Shutdown(cmd.instance, time.Second*3)

		return nil
	}
}
