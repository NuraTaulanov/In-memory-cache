In-memory cache implementation, Go package

go get -u -v  github.com/NuraTaulanov/In-memory-cache/
...
Get(key string) interface{}
Set(key string, value interface{})
Delete(key string)
...


	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}



