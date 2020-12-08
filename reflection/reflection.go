package reflection

import (
	"fmt"
	"reflect"
)

type WalkedFunc func(string)

func Walk(x interface{}, fn WalkedFunc) {

	v := reflect.ValueOf(x)
	// t := reflect.TypeOf(x)
	t := v.Type()
	p := reflect.ValueOf(&x)

	fmt.Println(p.Elem())
	// fmt.Println(v.Elem())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("%s => %v\n", t.Field(i).Name, v.Field(i))
		fmt.Printf("%s => %v\n\n", t.Field(i).Type, v.Field(i).Type())
	}
}
