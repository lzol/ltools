package util

import (
	"path/filepath"
	"strings"
	"github.com/pkg/errors"
	"os"
)

/*
	以import方式导入数据
	构造文件import_fileName，执行命令：sqlite3 dbName < import_fileName
*/
func ImportBigData(fileName string, dbName string, tableName string) (err error) {
	//判断文件是否为csv
	ext := filepath.Ext(fileName)
	if !strings.EqualFold(ext, ".csv") {
		return errors.New("文件格式不符")
	}
	//判断数据库文件是否存在
	f, err := os.Stat(dbName)
	if f.IsDir() || os.IsExist(err) {
		return errors.New("数据库文件不存在，或输入为目录")
	}
	params := []string{dbName}
	_, err = ExecCommand("sqlite3", params, false)

	return nil
}
