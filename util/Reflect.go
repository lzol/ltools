package util

import (
	"reflect"
	"errors"
)

/*
  根据类名，创建类实例
 */
func GetStructByName(structName string,structMap map[string]interface{})(ret interface{},err error){
	structType := structMap[structName]
	if structType == nil{
		return nil,errors.New("没有对应的类型名称")
	}
	t := reflect.ValueOf(structType).Type()
	ret = reflect.New(t).Elem().Interface()
	return ret,nil
}
