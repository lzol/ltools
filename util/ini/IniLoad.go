package common

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type IniFile struct {
	cfg *ini.File
	IniFile string
}

func NewIniFile(iniFile string)(*IniFile){
	t := new(IniFile)
	t.IniFile = iniFile
	cfg,err := ini.Load(t.IniFile)
	if err != nil {
		fmt.Println("Fail to read file: ", err)
	}
	t.cfg = cfg
	return t
}

func (t *IniFile)GetIntValue(section string,key string)(value int){
	v,err := t.cfg.Section(section).GetKey(key)
	if err!=nil{
		fmt.Printf("Fail to get value: ", section,key,err)
	}
	ret,err := v.Int()
	if err!=nil{
		fmt.Printf("Fail to get value: ", section,key,err)
	}
	return ret
}

func (t *IniFile)GetStringValue(section string,key string)(value string){
	v,err := t.cfg.Section(section).GetKey(key)
	if err!=nil{
		fmt.Printf("Fail to get value: %v%v%v", section,key,err)
	}
	ret := v.String()
	if err!=nil{
		fmt.Printf("Fail to get value: %v%v%v", section,key,err)
	}
	return ret
}
