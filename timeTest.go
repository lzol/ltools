package main

import (
	"fmt"
	"ltools/util"
	"time"
)

func main() {
	standTime := time.Now()
	actTime := standTime.Add(22*time.Second)
	fmt.Println(standTime,actTime)
	fmt.Println(actTime.Format("15:04:05"))

	r,_:=util.AddTime("15:04:05","18:00:00",22*time.Second)
	fmt.Println(r)
}
