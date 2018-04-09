package util

import (
	"path/filepath"
	"strings"
	"github.com/pkg/errors"
	"os"
	"io/ioutil"
	"fmt"
)

/*
	以import方式导入数据
	构造文件fileName_import，执行命令：sqlite3 dbName < fileName_import
*/
func ImportBigData(fileName string, dbName string, tableName string) (err error) {
	//判断文件是否为csv
	ext := filepath.Ext(fileName)
	if !strings.EqualFold(ext, ".csv") {
		return errors.New("文件格式不符")
	}
	//判断数据库文件是否存在
	f, err := os.Stat(dbName)
	fmt.Println(dbName)
	fmt.Println(f)
	fmt.Println(os.IsExist(err))
	if os.IsNotExist(err) || f.IsDir()  {
		return errors.New("数据库文件不存在，或输入为目录")
	}
	sqliteFileName := fileName+"_import"
	f,err = os.Stat(sqliteFileName)
	if os.IsExist(err){
		os.Remove(sqliteFileName)
	}
	file,err := os.Create(sqliteFileName)
	defer file.Close()
	sqliteCmdStr := ".mode csv\n"
	sqliteCmdStr += ".import "+fileName +" " + tableName+"\n"

	ioutil.WriteFile(sqliteFileName,[]byte(sqliteCmdStr),os.ModeAppend)
	params := []string{dbName,"<",sqliteFileName}
	fmt.Println(params)
	//line, err := ExecCommand("/usr/bin/sqlite3 "+dbName +" < "+ sqliteFileName, nil, true)
	line, err := ExecCommand("sqlite3", params, true)
	fmt.Println(line)
	return nil
}
