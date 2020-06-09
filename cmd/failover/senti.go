package main

import (
	"context"
	"fmt"
	"log"

	rediscli "github.com/go-redis/redis/v7"
)

type RedisMaster struct {
	IP   string
	PORT string
}

// GetRedisMaster getting Redis Master Ip and other information from the Sentinel
func GetRedisMaster(servicename string, mastername string) (RedisMaster, error) {
	ctx, _ := context.WithCancel(context.Background())
	senti := rediscli.NewSentinelClient(&rediscli.Options{
		Addr:     servicename,
		Password: "",
		DB:       0,
	})
	cmd := rediscli.NewSliceCmd("sentinel", "master", mastername)
	senti.ProcessContext(ctx, cmd)
	res, err := cmd.Result()
	if err != nil {
		log.Println("ERROR: ", err)
		return RedisMaster{}, err
	}
	return RedisMaster{
		PORT: fmt.Sprintf("%v", res[5]),
		IP:   fmt.Sprintf("%v", res[3]),
	}, nil
}
