# Friends Management
Friends Management API

## To generate APIs using Swagger

1. To install Swagger: `alias swagger='docker run --rm -e GOPATH=${GOPATH}:/go -v $(pwd):$(pwd) -w $(pwd) -u $(id -u):$(id -u) stratoscale/swagger:v1.0.21'`

2. To validate Swagger yml file: `swagger validate swagger/swagger.yaml`

3. To generate Go Server components and docs: `swagger generate server -A friends-management -f swagger/swagger.yaml`

## Installation
Using Docker Compose
To Setup the images:

* `$ docker-compose build`

To Start the development server:

* `$ docker-compose up`

To stop the development server:

* `$ docker-compose stop`

Stop Docker development server and remove containers:

* `$ docker-compose down`

## API Documentation
http://localhost:8888/api/v1/docs

## Dependancies/Technologies Used
* [Go](https://golang.org/) The Go Language.
* [MongoDB](https://www.mongodb.com/) MongoDB.
* [mgo](https://gopkg.in/mgo.v2) MongoDB driver for Go.
* [go-swagger](https://goswagger.io/) Swagger 2.0 for [Go].
* [Docker](https://www.docker.com/) Docker/Docker compose.