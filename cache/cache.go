package cache

type MemoryCache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
}

func New() *Cache {
	return &Cache{
		data: map[string]cacheItem{},
	}
}

func (c *Cache) Get(key string) interface{} {
	return c.data[key]
}
func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = cacheItem{
		value: value,
	}
}
func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
