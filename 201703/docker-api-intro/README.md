# GolangPhilly Talk

This talk was presented on March 21st, 2017 entitled "Introduction to the Docker Remote API with Go"

In this repo are the talk slides along with supporting code.

## Slides:

In order to view the slides you can use the golang present tool:

`present presentation.slide`

You can then view it at: http://localhost:3999/presentation.slide#1


## Demo Apps:

The demo applications can be found in the [*examples*](./examples) directory.

You'll need to get the Docker API packages:

`go get github.com/docker/docker/api/types/...`

`go get github.com/docker/docker/client/...`


Applications with a single main.go file can be run with:

`go run main.go`

Applications with more than one Go file will need to be built from the within its own directory with:

`go build .`

`./<app_dir_name>`

Applications with a *Makefile* are meant to be executed as a Docker container and are built and executed with:

`make docker`

`docker run -d robertojrojas/<app_name>`

#### NOTE: The demo apps are meant to demonstrate use of the Docker Engine API and do not represent best practices to be used in production applications.
