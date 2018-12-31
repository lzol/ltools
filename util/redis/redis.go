package redis

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*
 use github.com/astaxie/beego/cache/redis
*/

//type RedisConfig struct {
//	Key      string `json:"key"`      //key: Redis collection 的名称
//	Conn     string `json:"conn"`     //Redis 连接信息
//	DbNum    string    `json:"dbNum"`    //dbNum: 连接 Redis 时的 DB 编号. 默认是0.
//	Password string `json:"password"` // password: 用于连接有密码的 Redis 服务器.
//}

var redisConn cache.Cache

func InitByJson(jsonFile string) (err error) {
	f, err := os.Open(jsonFile)
	if err != nil {
		log.Println("init json file failed,", err)
		return  err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		if err != nil {
			log.Println("init json file failed,", err)
			return  err
		}
	}
	confStr := string(b)
	redisConn, err = cache.NewCache("redis", confStr)
	if err != nil {
		log.Println("init err:", err)
	}
	return err
}

func GetRedisConn()(cache.Cache){
	return redisConn
}

func GetString(key string)(value string){
	v := redisConn.Get(key)
	return string(v.([]byte))
}

func Put(key string,value interface{},duration time.Duration)(err error){
	return redisConn.Put(key,value,duration)
}

