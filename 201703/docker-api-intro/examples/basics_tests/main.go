package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"log"
	"os"
)

func main() {

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		cli.Close()
	}()

	args := filters.NewArgs()
	clo := types.ContainerListOptions{
		//All:true,
		Filters: args,
	}

	containerLister, err := New(context.Background(), cli, clo)
	if err != nil {
		log.Fatal(err)
	}
	displayContainers(containerLister, os.Stdout)
}
