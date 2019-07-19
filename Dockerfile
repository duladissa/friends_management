# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.11-alpine base image
FROM golang:1.11-alpine as build_base

# Set the Current Working Directory inside the container
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR $GOPATH/src/github.com/duladissa/friends_management

# Force the go compiler to use modules
ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the  go mod download 
# command will _ only_ be re-run when the go.mod or go.sum file change 
# (or when we add another docker instruction this line)
RUN go mod download

FROM build_base AS server_builder
# Here we copy the rest of the source code
COPY . .
# And compile the project
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -tags friends_management -ldflags '-w -extldflags "-static"' -o friends_management_server

FROM alpine:3.8
# This is needed for Mongo to work
RUN apk --no-cache add ca-certificates bash
WORKDIR /app/
COPY --from=server_builder /go/src/github.com/duladissa/friends_management/friends_management_server .
ENV PORT 8888
EXPOSE 8888
# Run the executable
CMD ["./friends_management_server"]