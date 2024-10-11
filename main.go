package main

import (
	"context"
)

type Foo struct{}

func (f *Foo) GetCacheVolumeId(ctx context.Context) (string, error) {
	cacheVolume := dag.CacheVolume("volume-name")
	_, err := dag.Container().
		From("alpine:latest").
		WithMountedCache("/foo", cacheVolume).
		WithExec([]string{"sh", "-c", "echo -n 'hello foo' > /foo/bar.txt"}).
		Sync(ctx)
	if err != nil {
		return "", err
	}

	id, err := cacheVolume.ID(ctx)
	return string(id), err
}

func (f *Foo) UseCacheVolumeAcrossModuleUsingId(ctx context.Context, id string) (string, error) {
	return dag.Bar().UseCacheVolumeID(ctx, id)
}
