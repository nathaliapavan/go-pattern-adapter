package main

import (
	"fmt"
	"go-pattern-adapter/infra/cache"
	"log"
)

func main() {
	var cacheService cache.Cache = cache.NewRedisAdapter()

	cacheKey := "myKey"
	err := cacheService.Set(cacheKey, "myValue")
	if err != nil {
		log.Fatalf("Error setting value in cache: %v", err)
	}

	value, err := cacheService.Get(cacheKey)
	if err != nil {
		log.Fatalf("Error getting value from cache: %v", err)
	}

	fmt.Println("Value obtained from cache:", value)

	err = cacheService.Delete(cacheKey)
	if err != nil {
		log.Fatalf("Error deleting value from cache: %v", err)
	}

	fmt.Println("Value deleted from cache successfully!!")
}
