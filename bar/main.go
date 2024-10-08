package main

import (
	"context"
)

type Bar struct{}

func (f *Bar) GetCacheVolumeId(ctx context.Context) (string, error) {
	id, err := dag.CacheVolume("volume-name-check-else").ID(ctx)
	return string(id), err
}
