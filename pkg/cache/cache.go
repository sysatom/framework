package cache

import (
	"fmt"
	"github.com/dgraph-io/ristretto/v2"
	"github.com/sysatom/framework/pkg/types"
)

func Example() {
	cache, err := ristretto.NewCache(&ristretto.Config[string, []*types.KV]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	defer cache.Close()

	// set a value with a cost of 1
	cache.Set("key", []*types.KV{{"Key": "key", "Value": "value"}}, 1)

	// wait for value to pass through buffers
	cache.Wait()

	// get value from cache
	value, found := cache.Get("key")
	if !found {
		panic("missing value")
	}
	fmt.Printf("value: %#v\n", value)

	// del value from cache
	cache.Del("key")
}
