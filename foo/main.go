package main

import (
	"context"
)

type Foo struct{}

func (f *Foo) GetCacheVolumeId(ctx context.Context) (string, error) {
	id, err := dag.CacheVolume("volume-name-check-else").ID(ctx)
	return string(id), err
}
