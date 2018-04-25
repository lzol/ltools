package util

import (
	"os"
	"path/filepath"
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