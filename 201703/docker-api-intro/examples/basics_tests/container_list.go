package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
)

type Warehouse struct {
	ctx      context.Context
	options  types.ContainerListOptions
	delegate func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
}

func (cl *Warehouse) ContainerList() ([]types.Container, error) {
	if cl.delegate != nil {
		return cl.delegate(cl.ctx, cl.options)
	}

	return nil, nil
}

func New(cntx context.Context, cli *client.Client, options types.ContainerListOptions) (*Warehouse, error) {
	return &Warehouse{
		ctx:     cntx,
		options: options,
		delegate: func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
			return cli.ContainerList(ctx, options)
		},
	}, nil
}

func displayContainers(cl *Warehouse, w io.Writer) error {
	containers, err := cl.ContainerList()
	if err != nil {
		return err
	}

	for _, container := range containers {
		fmt.Fprintf(w, "%s - %s\n", container.ID[:12], container.Names[0])
	}

	return nil
}
