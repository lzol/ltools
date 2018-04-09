package main

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"fmt"
)


func main(){
	//params := []string{"-ltr","-a"}
	//line,err := util.ExecCommand("ls",params,true)
	//fmt.Println(line,err)
	//err := util.ImportBigData("/Volumes/Share/tmp/outTest.csv","/Volumes/Share/tmp/testDB","TEST")
	//fmt.Println(err)

	srcDb, err := sql.Open("sqlite3", "/Volumes/Share/tmp/testDB")
	fmt.Println(err)
	result,err := srcDb.Exec("select * from test")
	fmt.Println(result,err)
	result,err = srcDb.Exec(".mode csv")
	fmt.Println(result,err)
	result,err = srcDb.Exec(".import /Volumes/Share/tmp/outTest.csv TEST")
	fmt.Println(result,err)
}
