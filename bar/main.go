package main

import (
	"context"
	"dagger/bar/internal/dagger"
)

type Bar struct{}

func (f *Bar) UseCacheVolumeId(ctx context.Context, id string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedCache("/bar", dag.LoadCacheVolumeFromID(dagger.CacheVolumeID(id))).
		WithExec([]string{"sh", "-c", "ls /bar/bar.txt"}).
		Stdout(ctx)
}

func (f *Bar) UseCacheVolumeName(ctx context.Context, name string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedCache("/bar", dag.CacheVolume(name)).
		WithExec([]string{"sh", "-c", "ls /bar/bar.txt"}).
		Stdout(ctx)
}
