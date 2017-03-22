package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"os/signal"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/filters"
)

func main() {

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	ctx, cancelFn := context.WithCancel(context.Background())

	args := filters.NewArgs()
	args.Add("event", "kill")
	args.Add("event", "die")
	eventsOption := types.EventsOptions{
		Filters:args,
	}
	eventsChan, errChan := cli.Events(ctx, eventsOption)


	fmt.Print("\u001B[2J\u001B[0;0f")
	fmt.Println("Waiting for Docker events....")
	fmt.Printf("%s%10s%10s\n", "CONTAINER", "SIGNAL", "ACTION")
	go func() {
		for {
			select {
			case event := <- eventsChan:
				fmt.Printf("%5s%10s%9s\n", event.Actor.Attributes["name"], event.Actor.Attributes["signal"], event.Action)
			}
		}
	}()


	// Wait for program to be terminated
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case err := <-errChan:
			if err != nil {
				fmt.Printf("Error while getting events: %s", err)
				time.Sleep(200 * time.Millisecond)
				os.Exit(1)
			}
		case <-signalChan:
			fmt.Println("\nBueno pues adios!")
			cancelFn()
		        time.Sleep(200 * time.Millisecond)
			os.Exit(0)
		}
	}
}