package main

import (
	"fmt"
	"ltools/util"
)

func main(){
	params := []string{"-ltr","-a"}
	line,err := util.ExecCommand("ls",params,true)
	fmt.Println(line,err)
	err = util.ImportBigData("/Volumes/Share/tmp/outTest.csv","/Volumes/Share/tmp/testDB","TEST")
	fmt.Println(err)
}
