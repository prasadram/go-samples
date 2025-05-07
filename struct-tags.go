package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID string `json:"id"`
}

func main() {
	u := User{}
	t := reflect.TypeOf(u)
	field := t.Field(0)

	fmt.Println("Field Name: ", field.Name)
	fmt.Println("Raw Tag: ", field.Tag)
	fmt.Println("JSON Tag: ", field.Tag.Get("json"))
}
