# Friends Management
Friends Management API

## To generate APIs using Swagger

1. To install Swagger: `alias swagger='docker run --rm -e GOPATH=${GOPATH}:/go -v $(pwd):$(pwd) -w $(pwd) -u $(id -u):$(id -u) stratoscale/swagger:v1.0.21'`

2. To validate Swagger yml file: `swagger validate swagger/swagger.yaml`

3. To generate Go Server components and docs: `swagger generate server -A friends-management -f swagger/swagger.yaml`