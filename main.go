package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var WELCOME_MESSAGE = "Hello word 12/08/2021\n"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	opts := redis.Options{
		Addr:     "redis-master.default.svc.cluster.local:6379",
		Password: "",
		DB:       0,
	}
	redisClient := redis.NewClient(&opts)
	err := redisClient.Set(ctx, "key", WELCOME_MESSAGE, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, val)
}
