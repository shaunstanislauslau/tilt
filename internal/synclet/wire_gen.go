// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package synclet

import (
	"context"
	"github.com/windmilleng/tilt/internal/docker"
	"github.com/windmilleng/tilt/internal/k8s"
)

// Injectors from wire.go:

func WireSynclet(ctx context.Context, env k8s.Env) (*Synclet, error) {
	dockerCli, err := docker.DefaultDockerClient(ctx, env)
	if err != nil {
		return nil, err
	}
	synclet := NewSynclet(dockerCli)
	return synclet, nil
}
