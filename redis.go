package main

import (
	"fmt"
	"regexp"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func dbInit() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func dbStatus() {
	res, err := client.Info("keyspace").Result()
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(".*keys=(\\d+),.*")
	match := re.FindStringSubmatch(res)
	fmt.Println("Database size: " + match[1] + " key(s)")
}

func dbCreate(k, v string) bool {
	if exists := dbGet(k); len(exists) > 0 {
		return false
	}
	return client.Set(k, v, 0).Err() == nil
}

func dbGet(k string) string {
	res, _ := client.Get(k).Result()
	return res
}

func redisInit() {
	dbInit()
	dbStatus()
}
