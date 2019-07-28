package main

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type Base struct {
	basename string
}
type Derive struct { // 内嵌 匿名成员
	Base
	adf int
}
type Derive1 struct { // 内嵌， 这种内嵌与上面内嵌有差异
	*Base
	adf int
}
type Derive2 struct { // 聚合
	base Base
	adf  int
}

var ErrNil = errors.New("redigo: nil returned")
func Sqrt(f float64) (z float64, err error) {
	if f < 0 {
		return 0, ErrNil
	}
	return 0,nil
}

func main() {
	//fmt.Printf("hello, world\n")
	//fmt.Println(cache.Name)
	_,err := Sqrt(-1)
	if err != nil{
		fmt.Print(err.Error())
	}
}

func statement() {
	fmt.Println(nil)
	var a int
	fmt.Println(a)
	var b float64
	fmt.Println(b)
	var c []byte
	fmt.Println(c)
	var d map[string]string
	fmt.Println(d)
	var e string
	fmt.Println(e)
	fmt.Println(e == "")
	fmt.Println(len(e) == 0)
	var f bool
	fmt.Println(f)
}

//func (api *LohCloudAPI) login() {
//	api.Header["fafa"] = "fadfadf"
//	fmt.Print(api.Header)
//}
//func s() {
//	type LohCloudAPI struct {
//		Header       map[string]string
//		MasterToken  []byte
//		SensorToken  []byte
//		//SandBoxToken []byte
//	}
//	api := LohCloudAPI{}
//	getType(api)
//
//	b := &LohCloudAPI{}
//	getType(b)
//
//	c := new(LohCloudAPI)
//	getType(c)
//	//api.login()
//
//}

func getType(a interface{}) {
	fmt.Print(reflect.TypeOf(a), reflect.ValueOf(a).Kind(), unsafe.Sizeof(a), "\t")
	fmt.Printf("%#v\n", a)
}
