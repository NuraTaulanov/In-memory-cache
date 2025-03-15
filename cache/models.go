package cache

import (
	"sync"
)

type cacheItem struct {
	value interface{}
}
type Cache struct {
	sync sync.RWMutex
	data map[string]cacheItem
}
