package main

import (
	"fmt"
<<<<<<< HEAD
	"reflect"
	"ltools/util"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	ttt(&t)

	fmt.Println("通过结构名称实例化对象")
	structMap := make(map[string]interface{})
	structMap["T"] = T{}

	a,_:= util.GetStructByName("T",structMap)
	b := a.(T)
	b.A = 100
	fmt.Println(b)

	structType := structMap["T"]
	if structType != nil{
		t := reflect.ValueOf(T{}).Type()
		ret := reflect.New(t).Elem().Interface()
		obj := ret.(T)
		obj.A = 200
		obj.B = "222"
		fmt.Println(obj)
	}


}

func ttt(t interface{}){
	tt := reflect.TypeOf(t)
	fmt.Printf("t type:%v\n", tt)
	ttp := reflect.TypeOf(&t)
	fmt.Printf("t type:%v\n", ttp)
	// 要设置t的值，需要传入t的地址，而不是t的拷贝。
	// reflect.ValueOf(&t)只是一个地址的值，不是settable, 通过.Elem()解引用获取t本身的reflect.Value
	s := reflect.ValueOf(t).Elem()
	fmt.Println(s)
	s.FieldByName("A").Set(reflect.ValueOf(111))

	fmt.Println(t)
}
=======
	"ltools/util"
)

func main(){
	params := []string{"-ltr","-a"}
	line,err := util.ExecCommand("ls",params,true)
	fmt.Println(line,err)
	err = util.ImportBigData("/Volumes/Share/tmp/outTest.csv","/Volumes/Share/tmp/testDB","TEST")
	fmt.Println(err)
}
>>>>>>> 3f9cd45432665067fc79eb4bccdfcff524ad9e0c
