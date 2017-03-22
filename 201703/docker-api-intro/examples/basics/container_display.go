package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"io"
)


func displayContainers(containers []types.Container, w io.Writer) error {
	for _, container := range containers {
		fmt.Fprintf(w, "%s - %s\n", container.ID[:12], container.Names[0])
	}

	return nil
}

