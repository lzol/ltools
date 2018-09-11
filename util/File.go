package util

import (
	. "io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateNewFile(fileName string)(file *os.File,err error){
	_, err = os.Stat(fileName)
	if os.IsExist(err){
		os.Remove(fileName)
	}
	file,err = os.Create(fileName)
	if err!=nil{
		return nil,err
	}
	return file,nil
}

/*
	获取文件名称，不含路径
*/
func GetFileName(fileWholePath string)( fileName string){
	return filepath.Base(fileWholePath)
}

/*
	获取文件路径，不含名称
*/
func GetFilePath(fileWholePath string)( filePath string){
	return filepath.Dir(fileWholePath)
}

/*
	获取文件路径，不含名称
*/
func GetFileExt(fileWholePath string)( fileExt string){
	return filepath.Ext(fileWholePath)
}

/*
	获取文件全路径，含名称
*/
func GetFileWholePath(fileName string)( fileWholePath string,err error){
	fileWholePath,err = filepath.Abs(fileName)
	return fileWholePath,err
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}