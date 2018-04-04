package test

import (
	"testing"
	"container/list"
	"fmt"
)

func TestStringSplit(t *testing.T){

}

type Tt struct {
	Name string
	Age int
}

func TestList(t *testing.T){
	list1 := list.New()
	//list1.PushBack(1)
	//list1.PushBack(2)
	map1 := map[string]string{
		"a":"a",
	}
	map2 := map[string]Tt{
		"a":Tt{"name",16},
	}
	list1.PushBack(map1)
	list1.PushBack(map2)
	for e := list1.Front(); e != nil; e = e.Next() {
		tmp := e.Value
		fmt.Println(tmp) //输出list的值,无内容
	}
}
