package util

import "os"

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
