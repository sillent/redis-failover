package main

import (
	"context"
	"fmt"
	"log"

	rediscli "github.com/go-redis/redis/v7"
)

// RedisMaster struct contained IP and Port field
type RedisMaster struct {
	IP   string
	PORT string
}

// MasterAsStr return string like "127.0.0.1:3333"
func (r RedisMaster) MasterAsStr() string {
	return fmt.Sprintf("%v:%v", r.IP, r.PORT)
}

// getRedisMaster getting Redis Master Ip and other information from the Sentinel
func getRedisMaster(servicename string, mastername string) (RedisMaster, error) {
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
