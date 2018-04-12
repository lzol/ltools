package main

import (
	_ "database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	_ "ltools/util"
	"os/exec"
	"bytes"
	"ltools/util"
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
	//err := util.ImportBigData("/usr/bin/sqlite3","/Volumes/Share/GOPATH/src/duizhang/csv/aceve.csv", "/Volumes/Share/GOPATH/src/duizhang/data/duizhang", "t_aceve")
	result,err := util.ImportBigData("D://Program//sqlite-tools-win32-x86-3230100//sqlite3","e://GOPATH//src//duizhang//csv//aceve.csv", "e://GOPATH//src//duizhang//data//duizhang", "t_aceve")
	fmt.Println(result,err)
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
