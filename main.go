package main

import (
	"bytes"
	_ "database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/streadway/amqp"
	_ "ltools/util"
	"ltools/util/amqp"
	"os/exec"
	_ "strings"
)

func main() {
	//params := []string{"-ltr","-a"}
	//line,err := util.ExecCommand("ls",params,true)
	//fmt.Println(line,err)
	//err := util.ImportBigData("/Volumes/Share/tmp/outTest.csv","/Volumes/Share/tmp/testDB","TEST")
	//fmt.Println(err)

	//srcDb, err := sql.Open("sqlite3", "/Volumes/Share/tmp/testDB")
	//fmt.Println(err)
	//result,err := srcDb.Exec("select * from test")
	//fmt.Println(result,err)
	//result,err = srcDb.Exec(".mode csv")
	//fmt.Println(result,err)
	//result,err = srcDb.Exec(".import /Volumes/Share/tmp/outTest.csv TEST")
	//fmt.Println(result,err)

	//params := []string{"/Volumes/Share/tmp/test.sh"}
	//line,err := util.ExecCommand("/bin/sh",params,true)
	//line,err := util.ExecCommand("e:/tmp/test.bat",nil,true)
	//fmt.Println(line,err)
	//fmt.Println(runtime.GOOS)
	//exec_shell("./test.sh")
	//result,err := util.ImportBigData("/Users/liuz/Library/Android/sdk/platform-tools/sqlite3","/Volumes/Share/GOPATH/src/duizhang/csv/aceve.csv", "/Volumes/Share/GOPATH/src/duizhang/data/duizhang", "t_aceve")
	//result,err := util.ImportBigData("D://Program//sqlite-tools-win32-x86-3230000//sqlite3","e://GOPATH//src//duizhang//csv//aceve.csv", "e://GOPATH//src//duizhang//data//duizhang", "t_aceve")
	//fmt.Printf("result:[%s] err:[%v]\n",result,err)
	//fmt.Println("go"=="go")
	//fmt.Println("GO"=="go")

	//fmt.Println(strings.Compare("GO","go"))
	//fmt.Println(strings.Compare("go","go"))

	//fmt.Println(strings.EqualFold("GO","go"))

	//result,err = util.ExecCommand("cmd",[]string{"/C","dir"},true)
	//fmt.Println(util.Encode(result,util.GBK,util.UTF8),err)
	//s := "0305-ACEVE-20180416"
	//s1 := s[len(s)-8:len(s)]
	//fmt.Println(s1)
	//uuid,_ := util.GetUUID()
	//fmt.Println(uuid)
	//uuid,_ = util.Get32UUID()
	//fmt.Println(uuid)
	exchange := new(util.Exchange)
	exchange.Name = "test"
	exchange.Type = util.EXCHANGE_TYPE_DIRECT
	exchange.Durable = true
	queue := new(util.Queue)
	queue.Name = "test"
	queue.Durable = true
	amqp := new(util.Amqp)
	amqp.URI = "amqp://guest:guest@localhost:5672/"
	amqp.Queue = *queue
	amqp.Exchange = *exchange
	defer amqp.Close()
	amqp.Public("sfdljdsfldsjfls")
	fmt.Println("发布队列消息：sfdljdsfldsjfls")
	amqp.Consume(OnReceive)

}

func OnReceive(d amqp.Delivery){
	fmt.Println("收到队列消息：",string(d.Body))
}

func exec_shell(s string) {
	cmd := exec.Command("/bin/sh", s)
	var out bytes.Buffer
	fmt.Println(cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("%s", out.String())
	if err != nil {
		fmt.Println("err", err)
	}

}
