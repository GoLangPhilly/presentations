package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
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
		fmt.Printf("Usage: %s <Network Name> <CDIR>", os.Args[0])
		os.Exit(1)
	}

	networkName := os.Args[1]
	subnet      := os.Args[2]

	ipamConfigs := []network.IPAMConfig{
		{
		   Subnet:subnet,
		},
	}

	ipam := &network.IPAM{
		Config:ipamConfigs,
	}
	options := types.NetworkCreate{
		IPAM:ipam,
	}

	ncr, err := cli.NetworkCreate(context.Background(), networkName, options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Network ID: %s\n", ncr.ID)
}