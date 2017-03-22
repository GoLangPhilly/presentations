package main

import (
	"bytes"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"testing"
)

func TestDisplayContainers(t *testing.T) {

	cases := []struct {
		delegate    func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
		expectedStr string
		expectedErr error
	}{
		{
			delegate: func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
				containers := []types.Container{
					{
						ID:    "bec9a3a8e3adbec9a3a8e3ad",
						Names: []string{"Container1"},
					},
					{
						ID:    "3bea056a69963bea056a6996",
						Names: []string{"Container2"},
					},
				}
				return containers, nil

			},
			expectedStr: "bec9a3a8e3ad - Container1\n3bea056a6996 - Container2\n",
		},
		{
			delegate: func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
				return nil, errors.New("no containers")
			},
			expectedErr: errors.New("no containers"),
		},
	}

	for _, containerCase := range cases {
		mock := &Warehouse{
			delegate: containerCase.delegate,
		}

		var bf bytes.Buffer
		err := displayContainers(mock, &bf)

		if err != nil && err.Error() != containerCase.expectedErr.Error() {
			t.Error("Errors do not match")
		}

		if len(containerCase.expectedStr) > 0 && containerCase.expectedStr != bf.String() {
			t.Errorf("expected [%s][%s] instead\n", containerCase.expectedStr, bf.String())
		}
	}

}
