package main

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"os"
	"time"
	"fmt"
	"syscall"
	"os/signal"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <Container ID/Name>", os.Args[0])
		os.Exit(1)
	}
	containerIdOrName := os.Args[1]

	ctx, cancelFn := context.WithCancel(context.Background())
	response, err := cli.ContainerStats(ctx, containerIdOrName, true)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	go func() {
		dec := json.NewDecoder(response.Body)
		for {
			var containerStats *types.StatsJSON
			if err := dec.Decode(&containerStats); err != nil {
				dec = json.NewDecoder(io.MultiReader(dec.Buffered(), response.Body))
				if err == io.EOF {
					break
				}
				time.Sleep(100 * time.Millisecond)
				continue
			}
			displayContainerStats(containerStats)
		}
	}()


	// Wait for program to be terminated
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			fmt.Println("\nBueno pues adios!")
			cancelFn()
			time.Sleep(200 * time.Millisecond)
			os.Exit(0)
		}
	}
}
