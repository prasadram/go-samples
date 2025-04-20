package main

import (
	"fmt"
	"reflect"
    )
   
type MyStruct struct {
	SomeField int
}

func (ms *MyStruct) SomeMethod() int {
	return ms.SomeField
}

func GetStructMethodInfo(funcPointer interface{}) {
	funcType := reflect.TypeOf(funcPointer)
	fmt.Printf("Function Type: %v\n", funcType)
}

func GetStructInfo(structPointer interface{}) {
	structType := reflect.TypeOf(structPointer)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	fmt.Printf("Struct name: %s\n", structType.Name())
	
	method, found := reflect.TypeOf(structPointer).MethodByName("SomeMethod")
	if found {
		fmt.Printf("Method name: %s\n", method.Name)
		fmt.Printf("Method type: %v\n", method.Type)
	}
}
	
func main() {
	ms := &MyStruct{SomeField: 45}
	GetStructMethodInfo(ms.SomeMethod)
	GetStructInfo(ms)
}
