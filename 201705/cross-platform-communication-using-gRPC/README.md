# GolangPhilly Talk

This talk was presented on April 11th, 2017 entitled "Cross-Platform Communication Using gRPC"


In this repo are the talk slides along with supporting code.

## Slides:

In order to view the slides you can use the golang present tool:

`present presentation.slide`

You can then view it at: http://localhost:3999/presentation.slide#1


## Demo Apps:

The demo applications can be found in [*examples*](./examples).

You'll need to get the gRPC packages:

`go get google.golang.org/grpc`

`go get google.golang.org/grpc/credentials`

`go get google.golang.org/grpc/metadata`


Applications with a single main.go file can be run with:

`go run main.go`

Applications with more than one Go file will need to be built from the within its own directory with:

`go build .`

`./<app_dir_name>`

#### NOTE: The demo apps are meant to demonstrate use of gRPC and do not represent best practices to be used in production applications.
