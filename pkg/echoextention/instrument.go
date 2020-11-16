package echoextention

import (
	"github.com/google/uuid"
	echoLog "github.com/labstack/gommon/log"

	//"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

//Root Level (After router)
//The following built-in middleware should be registered at this level:
//BodyLimit
//Logger
//Gzip
//Recover
//ServerHeader middleware adds a `Server` header to the response.
func RegisterGlobalMiddlewares(e *echo.Echo) {

	e.Use(middleware.BodyLimit("1M"))
	addCollerationId(e)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: Myskipper,
		Level:   -1,
	}))
	//e.Use(HookGateLoggerWithConfig(GateLoggerConfig{
	//	IncludeRequestBodies:  false,
	//	IncludeResponseBodies: false,
	//	Skipper:               Myskipper,
	//}))
	e.Use(RecoverWithConfig(RecoverConfig{
		Skipper:           Myskipper,
		StackSize:         4 << 10,
		DisableStackAll:   false,
		DisablePrintStack: false,
		LogLevel:          echoLog.INFO,
		statusCodeMapping: nil,
	}))
}

func Myskipper(context echo.Context) bool {
	if strings.HasPrefix(context.Path(), "/status") ||
		strings.HasPrefix(context.Path(), "/swagger") ||
		strings.HasPrefix(context.Path(), "/metrics") {
		return true
	}

	return false
}
func addCollerationId(e *echo.Echo) {
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			uid, _ := uuid.NewRandom()
			return uid.String()
		},
	}))
}
