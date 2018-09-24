package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	redis2 "ltools/util/redis"
	"time"
)

func main() {

	err := redis2.InitByJson("redis.json")
	if err!=nil{
		fmt.Println(err)
	}
	bm := redis2.GetRedisConn()
	err = bm.Put("sex1","male111",23*time.Second)
	if err!=nil{
		fmt.Println(err)
	}
	for {
			v := bm.Get("sex1")
			fmt.Println(redis.String(v, err))
			time.Sleep(3 * time.Second)
	}

}

