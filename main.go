package main

import (
	"log"
	"os"

	"github.com/gocraft/work"
	"github.com/gofiber/fiber/v2"
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

var enqueuer = work.NewEnqueuer("APICaller", redisPool)

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		body := new(struct {
			Seconds int    `json:"seconds"`
			URL     string `json:"url"`
			Payload string `json:"payload"`
			Type    string `json:"type"`
		})
		if err := c.BodyParser(body); err != nil {
			return err
		}
		_, err := enqueuer.EnqueueIn("FiberAPI", int64(body.Seconds), work.Q{"url": body.URL, "payload": body.Payload, "type": body.Type})
		if err != nil {
			log.Fatal(err)
		}
		return c.SendString("success")
	})

	app.Listen(":3000")
}
