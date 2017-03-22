package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
	"log"
	"os"
	"github.com/docker/docker/api/types/network"
)


func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <Network ID/Name> <Container ID/Name>\n", os.Args[0])
		os.Exit(1)
	}

	networkName    := os.Args[1]
	containerName  := os.Args[2]

	config := &network.EndpointSettings{}

	err = cli.NetworkConnect(context.Background(),networkName, containerName, config )
        if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connected %s to %s\n", networkName, containerName)
}