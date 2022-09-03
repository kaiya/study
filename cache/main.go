package main

import (
	"os"

	"gitlab.momoso.com/cm/kit/third_party/lg"
	"gitlab.momoso.com/mms2/utils/cache"
)

func main() {
	dir := os.TempDir()

	defer func() {
		os.RemoveAll(dir)
	}()

	diskCache, err := cache.NewDiskCache(dir)
	lg.PanicError(err, "create diskCache error")

	cacheSet := make(map[string]bool)
	diskCache.GetOrCreate("newCacheKey", func() (interface{}, error) {
		cacheSet["setKey"] = true
		return &cacheSet, nil
	}, &cacheSet)

	for key, value := range cacheSet {
		lg.Infof("key:%s, value:%v", key, value)
	}
}
