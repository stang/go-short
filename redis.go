package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func dbInit() {
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	client = redis.NewClient(opt)
}

func dbStatus() {
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong + " - connection established with redis server")

	res, err := client.Info("keyspace").Result()
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(".*keys=(\\d+),.*")
	match := re.FindStringSubmatch(res)
	if len(match) > 1 {
		fmt.Println("Database size: " + match[1] + " key(s)")
	}
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
