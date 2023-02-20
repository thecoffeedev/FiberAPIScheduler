package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-resty/resty/v2"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", os.Getenv("REDIS_HOST") + ":6379", redis.DialPassword(os.Getenv("REDIS_PASSWORD")))
	},
}

type Context struct {
	Seconds int    `json:"seconds"`
	URL     string `json:"url"`
	Payload string `json:"payload"`
	Type    string `json:"type"`
}

func main() {
	pool := work.NewWorkerPool(Context{}, 10, "APICaller", redisPool)
	pool.Job("FiberAPI", (*Context).CallAPI)
	pool.Start()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	pool.Stop()
}

func (c *Context) CallAPI(job *work.Job) error {
	client := resty.New()
	URL := job.ArgString("url")
	Payload := job.ArgString("payload")
	Type := job.ArgString("type")
	fmt.Println(URL)
	fmt.Println(Payload)
	fmt.Println(Type)
	if Type == "POST" {
		client.SetHeader("Content-Type", "application/json")
		Payload = strings.ReplaceAll(Payload, "'", "\"")
		fmt.Println(Payload)
		resp, err := client.R().SetBody(Payload).Post(URL)
		fmt.Println(resp)
		if err != nil {
			fmt.Println(err)
		}
	} else if Type == "GET" {
		resp, err := client.R().Get(URL)
		fmt.Println(resp)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
