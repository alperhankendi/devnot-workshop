package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"time"
)

func BuildCommandHandler(instance *echo.Echo, s *resource) {

	g := instance.Group("api/v1")

	g.POST("/v3", s.CreateV3, durationMethod)
	g.POST("/v2", s.CreateV2, durationMethod)
	g.POST("/v1", s.CreateV1, durationMethod)
}
func BuildQueryHandler(instance *echo.Echo, s *resource) {

	g := instance.Group("api/v1")

	g.GET("/:id", s.GetV1, durationMethod)

}

func durationMethod(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		start := time.Now()
		var err error
		if err = next(c); err != nil {
			c.Error(err)
		}
		stop := time.Now()
		go fmt.Printf("took: %s\n", stop.Sub(start).String())
		return nil
	}
}
