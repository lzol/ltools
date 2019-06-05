package util

import (
	"ltools/util/redis"
	"time"
)

/*获取redis中的流水号，如果redis对应的key不存在，则先创建，过期时间为一天
  流水号规则为YYMMDDHHMMSS+序号，序号长度为len，不足前补0
 */
func GetFlowNoByKey(key string,len int,duration time.Duration)(ret string,err error){
	if !redis.Exist(key){
		redis.Put(key,1,duration)
	}
	v,err := redis.GetAndIncr(key)
	if err!=nil {
		return ret,err
	}
	no,err := StringFill(v,0,"0",len)
	if err!=nil{
		return ret,err
	}
	date := GetDateString("20060102150405")
	return date+no,err
}
