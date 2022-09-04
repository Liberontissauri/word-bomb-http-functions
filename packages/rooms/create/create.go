package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type player struct{
	id string
	authorization_key string
	username string
	owner bool
}

func Main(args map[string]interface{}) map[string]interface{} {
	redis_option, _ := redis.ParseURL(os.Getenv("DATABASE_URL"))
	rdb := redis.NewClient(redis_option)
	room_template := make(map[string]interface{})
	room_template["id"] = args["id"].(string)
	room_template["players"] = make([]player, 0)
	room_template["creation_time"] = time.Now().Unix()
	room_template["bomb_start_time"] = 0
	room_template["words"] = make([]string, 0)
	room_template["language"] = args["language"].(string)
	room_template["current_word"] = ""
	room_template["started"] = false
	room_template["finished"] = false

	rdb.Set(context.Background(), "room:"+args["room_id"].(string), room_template, 0)
	msg := make(map[string]interface{})
	msg["body"] = "Room created"
	return msg
}
