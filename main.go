package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var WELCOME_MESSAGE = "Hello word 12/08/2021\n"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world")
	// checkRedis(w)
	// checkPubSub(w)
}

func checkRedis(w http.ResponseWriter) {
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
	fmt.Fprint(w, "Value from redis: "+val)
}

func checkPubSub(w http.ResponseWriter) {
	os.Setenv("PUBSUB_EMULATOR_HOST", "pub-sub-emulator.default.svc.cluster.local:8085")
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "my-project")
	if err != nil {
		panic(err)
	}
	topic, err := client.CreateTopic(ctx, "my-topic")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, "Topic: "+topic.String())
	err = topic.Delete(ctx)
	if err != nil {
		panic(err)
	}
}
