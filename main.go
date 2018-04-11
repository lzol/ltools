package main

import (
	_"database/sql"
	_"github.com/mattn/go-sqlite3"
	"fmt"
	_"ltools/util"
	"os/exec"
	"bytes"
	"ltools/util"
	"runtime"
)


func main(){
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
	line,err := util.ExecCommand("e:/tmp/test.bat",nil,true)
	fmt.Println(line,err)
	fmt.Println(runtime.GOOS)
	//exec_shell("./test.sh")
}
func exec_shell(s string) {
	cmd := exec.Command("/bin/sh", s)
	var out bytes.Buffer
	fmt.Println(cmd)
	cmd.Stdout = &out
	err := cmd.Run()
	fmt.Printf("%s", out.String())
	if err != nil {
		fmt.Println("err",err)
	}

}
