/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/alperhankendi/devnot-workshop/pkg/echoextention"
	"github.com/alperhankendi/devnot-workshop/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

type queryCommand struct {
	instance *echo.Echo
	command  *cobra.Command
	Port     string
}

// queryCmd represents the query command
var queryCmd = &queryCommand{
	command: &cobra.Command{
		Use:   "query",
		Short: "Query Service",
	},
	Port: "5000",
}

func init() {
	rootCmd.AddCommand(queryCmd.command)
	queryCmd.command.Flags().StringVarP(&queryCmd.Port, "port", "p", "5000", "Service Port")

	queryCmd.instance = echo.New()
	queryCmd.instance.Debug = false
	queryCmd.instance.HidePort = true
	queryCmd.instance.HideBanner = true
	queryCmd.instance.Logger = log.SetupLogger()
	echoextention.RegisterGlobalMiddlewares(queryCmd.instance)
	queryCmd.command.RunE = func(cmd *cobra.Command, args []string) error {

		registerHandler(queryCmd.instance)
		log.Logger.Infof("Service is starting. Service port:%s", queryCmd.Port)
		go func() {
			if err := queryCmd.instance.Start(fmt.Sprintf(":%s", queryCmd.Port)); err != nil {
				log.Logger.Fatalf("Failed to shutting down the server. Error :%v", err)
			}
		}()
		echoextention.Shutdown(queryCmd.instance, time.Second*3)

		return nil
	}
}

func registerHandler(e *echo.Echo) {

	g := e.Group("v1")

	g.GET("/", func(context echo.Context) error {

		return context.String(http.StatusOK, string(mediumFixture))
	})
	g.GET("/1", func(context echo.Context) error {

		return context.JSON(http.StatusOK, string(mediumFixture))
	})

}

// Reponse from Size: 2.4kb
var mediumFixture []byte = []byte(`{
  "person": {
    "id": "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "name": {
      "fullName": "Leonid Bugaev",
      "givenName": "Leonid",
      "familyName": "Bugaev"
    },
    "email": "leonsbox@gmail.com",
    "gender": "male",
    "location": "Saint Petersburg, Saint Petersburg, RU",
    "geo": {
      "city": "Saint Petersburg",
      "state": "Saint Petersburg",
      "country": "Russia",
      "lat": 59.9342802,
      "lng": 30.3350986
    },
    "bio": "Senior engineer at Granify.com",
    "site": "http://flickfaver.com",
    "avatar": "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "employment": {
      "name": "www.latera.ru",
      "title": "Software Engineer",
      "domain": "gmail.com"
    },
    "facebook": {
      "handle": "leonid.bugaev"
    },
    "github": {
      "handle": "buger",
      "id": 14009,
      "avatar": "https://avatars.githubusercontent.com/u/14009?v=3",
      "company": "Granify",
      "blog": "http://leonsbox.com",
      "followers": 95,
      "following": 10
    },
    "twitter": {
      "handle": "flickfaver",
      "id": 77004410,
      "bio": null,
      "followers": 2,
      "following": 1,
      "statuses": 5,
      "favorites": 0,
      "location": "",
      "site": "http://flickfaver.com",
      "avatar": null
    },
    "linkedin": {
      "handle": "in/leonidbugaev"
    },
    "googleplus": {
      "handle": null
    },
    "angellist": {
      "handle": "leonid-bugaev",
      "id": 61541,
      "bio": "Senior engineer at Granify.com",
      "blog": "http://buger.github.com",
      "site": "http://buger.github.com",
      "followers": 41,
      "avatar": "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
    },
    "klout": {
      "handle": null,
      "score": null
    },
    "foursquare": {
      "handle": null
    },
    "aboutme": {
      "handle": "leonid.bugaev",
      "bio": null,
      "avatar": null
    },
    "gravatar": {
      "handle": "buger",
      "urls": [
      ],
      "avatar": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
      "avatars": [
        {
          "url": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
          "type": "thumbnail"
        }
      ]
    },
    "fuzzy": false
  },
  "company": null
}`)
