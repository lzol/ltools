package util

import (
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/*
	以import方式导入数据
	构造文件fileName_import，执行命令：sqlite3 dbName < fileName_import
*/
func ImportBigData(sqlitePath string, fileName string, dbName string, tableName string) (result string, err error) {
	//判断文件是否为csv
	ext := filepath.Ext(fileName)
	if !strings.EqualFold(ext, ".csv") {
		return result, errors.New("文件格式不符")
	}
	//判断数据库文件是否存在
	f, err := os.Stat(dbName)
	if os.IsNotExist(err) || f.IsDir() {
		return result, errors.New("数据库文件不存在，或输入为目录")
	}
	sqliteFileName, err := createSqliteFile(fileName, tableName)
	if err != nil {
		return result, err
	}
	cmdFile, err := createCmd(sqlitePath, sqliteFileName, dbName)
	if runtime.GOOS == "windows" {
		//params:= []string{"/C",cmdFile}
		params := []string{"/C",sqlitePath,
			dbName,
			"<",sqliteFileName,
		}

		result, err := ExecCommand("cmd", params, true)
		//fmt.Println(result,err)
		if err != nil {
			return result, err
		}

	} else {
		params := []string{cmdFile}
		result, err := ExecCommand("/bin/sh", params, true)
		//params := []string{"-c",sqlitePath,
		//	dbName,
		//	"<", sqliteFileName,
		//}
		//result, err := ExecCommand("/bin/bash", params, true)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}

func createSqliteFile(fileName string, tableName string) (sqliteFileName string, err error) {
	sqliteFileName = fileName + "_import"
	_, err = os.Stat(sqliteFileName)
	if os.IsExist(err) {
		os.Remove(sqliteFileName)
	}
	file, err := os.Create(sqliteFileName)
	if err != nil {
		return sqliteFileName, errors.New("创建sqlite导入文件失败")
	}
	defer file.Close()
	sqliteCmdStr := ".mode csv\n"
	wholeFileName, _ := filepath.Abs(fileName)
	if runtime.GOOS == "windows" {
		wholeFileName = strings.Replace(wholeFileName, "\\", "//", -1)
	}
	sqliteCmdStr += ".import " + wholeFileName + " " + tableName + "\n"

	ioutil.WriteFile(sqliteFileName, []byte(sqliteCmdStr), os.ModeAppend)
	return sqliteFileName, nil
}
/*
	mac环境直接执行命令，无论原表存在与否，都会重新建表（/bin/bash -c sqlite3 db < import），导致无法正常导入
*/
func createCmd(sqlitePath, sqliteFileName, dbName string) (cmdFileName string, err error) {
	ext := ".sh"
	head := "#/usr/bin/sh"
	if runtime.GOOS == "windows" {
		ext = ".bat"
		head = ""
	}
	filepath.Base(sqliteFileName)
	cmdFileName = filepath.Dir(sqliteFileName) + "/import" + ext
	f, err := CreateNewFile(cmdFileName)
	defer f.Close()
	if err != nil {
		return cmdFileName, err
	}

	wholeDbName, _ := filepath.Abs(dbName)
	wholsSqliteFileName, err := filepath.Abs(sqliteFileName)
	io.WriteString(f, head+"\n")
	_, err = io.WriteString(f, sqlitePath+" "+wholeDbName+" < "+wholsSqliteFileName+"\n")
	if err != nil {
		return cmdFileName, err
	}
	return cmdFileName, nil
}
