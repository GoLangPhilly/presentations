package main

import (
	"log"
	"os"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types"
	"context"
	"github.com/docker/docker/client"
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
	filter := "web=nginx"
	args.Add("label", filter)
	clo := types.ContainerListOptions {
		//All:true,
		Filters:args,
	}

	containers, err := cli.ContainerList(context.Background(), clo)
	if err != nil {
		log.Fatal(err)
	}
	displayContainers(containers, os.Stdout)
}