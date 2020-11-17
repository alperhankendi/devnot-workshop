### Developer summit 2020@workshop
#### tools & libraries

- [Echo web framework](https://echo.labstack.com/)
- [Cobra](https://cobra.dev/) A Commander for modern Go CLI interactions 
- [K6](https://k6.io/docs/getting-started/installation)
- [Golang mongo driver](https://github.com/mongodb/mongo-go-driver)
- [Docker](https://www.docker.com) for mongo

#### struct

![Doc](docs/api.png)

#### folder structure

```
❯ tree
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── build
│   ├── deployment.yaml
│   └── service.yaml
├── cmd
│   ├── command.go
│   ├── query.go
│   └── root.go
├── docs
│   ├── api.png
│   └── api.puml
├── go.mod
├── internal
│   └── movies
│       ├── controller
│       │   ├── dispatcher.go
│       │   ├── endpoint.go
│       │   └── transport.go
│       ├── mongo
│       │   └── mongo_repository.go
│       ├── service.go
│       └── types.go
├── load-post.js
├── load.js
├── main.go
└── pkg
    ├── echoextention
    │   ├── gate_logger.go
    │   ├── healthcheck
    │   │   ├── healthcheck.go
    │   │   ├── healtycheck_test.go
    │   │   └── mongoChecker.go
    │   ├── instrument.go
    │   ├── recovery.go
    │   └── shutdown.go
    ├── errors
    │   ├── error.go
    │   ├── error_definition.go
    │   ├── errors.go
    │   └── statusCode.go
    ├── httpclient
    │   └── client.go
    ├── log
    │   └── logrus.go
    ├── mongoextentions
    │   ├── errors.go
    │   ├── mongo.go
    │   └── mongo_test.go
    └── utils
        └── mytime.go
```

#### build

```
$ git clone https://github.com/alperhankendi/devnot-workshop.git
$ git checkout api-final
$ cd devnot-workshop
$ make build
$ ./devnot-workshop
Application is starting. Version:Devnote workshop application

Usage:
  devnot-workshop [flags]
  devnot-workshop [command]

Available Commands:
  command     Command Service
  help        Help about any command
  query       Query Service

Flags:
  -c, --conn string     database connection string (default "mongodb://root:example@localhost:27017")
  -d, --dbname string   database name (default "imdb")
  -h, --help            help for devnot-workshop
  -p, --port string     Service Port (default "5001")

Use "devnot-workshop [command] --help" for more information about a command.
```

##### setup stuff

````
$ make dev-env-stop
docker stop demo-mongo demo-mongo-express
demo-mongo
demo-mongo-express

$ make dev-env-start
docker run --rm --name demo-mongo \
                -e MONGO_INITDB_ROOT_USERNAME=root -e  MONGO_INITDB_ROOT_PASSWORD=example -d \
                -p 27017:27017 mongo
f91e89e3432a21de34ae153a96d5a7f4b7bcbc2be7d38c8c8b354fcf7fbe8baf
docker run --rm --name demo-mongo-express \
                -e ME_CONFIG_MONGODB_SERVER=172.17.0.2 -e ME_CONFIG_MONGODB_PORT=27017 \
                -e ME_CONFIG_MONGODB_ADMINUSERNAME=root -e ME_CONFIG_MONGODB_ADMINPASSWORD=example -d \
                -p 8081:8081 mongo-express
d855275e34f666b46ad597c9471d9f9df39822a4e450549bd85fc199468dce23
 
````

##### run services
Command Service
```
./devnot-workshop api command -p 5001 -c mongodb://root:example@localhost:27017 -d imdb
```
Query Service
```
./devnot-workshop api query -p 5000 -c mongodb://root:example@localhost:27017 -d imdb
```
#### tests 

post some data
```
k6 run --duration 60s --vus 150 load-post.js 
```

get some data
```
k6 run --duration 60s --vus 150 load-get.js -e ID=7lf1ly4khsnhzfqsntec

running (1m00.1s), 000/150 VUs, 3126018 complete and 0 interrupted iterations
default ✓ [======================================] 150 VUs  1m0s
    ✗ status is 200
     ↳  99% — ✓ 3125731 / ✗ 287
```

