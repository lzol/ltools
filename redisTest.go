package main

import (
	"fmt"
	"ltools/util"
	redis2 "ltools/util/redis"
	"time"
)

func main() {

	err := redis2.InitByJson("redis.json")
	if err != nil {
		fmt.Println(err)
	}
	//bm := redis2.GetRedisConn()
	//err = bm.Put("sex2","msl",23*time.Second)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//for {
	//		v := bm.Get("sex1")
	//		fmt.Println(redis.String(v, err))
	//		time.Sleep(3 * time.Second)
	//}
	for i := 0; i < 10000; i++ {
		i, err := util.GetFlowNoByKey("aaa", 8,3*time.Second)
		fmt.Println(i, err)
	}
}
