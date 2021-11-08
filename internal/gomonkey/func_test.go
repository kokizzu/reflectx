package gomonkey

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMakeFunc(t *testing.T) {
	typ := reflect.FuncOf([]reflect.Type{reflect.TypeOf(0)}, []reflect.Type{reflect.TypeOf("")}, false)
	fn := MakeFunc(func() {}, typ, func(args []reflect.Value) []reflect.Value {
		r := fmt.Sprintf("str:%v", args[0].Int())
		return []reflect.Value{reflect.ValueOf(r)}
	})
	v := fn.Interface().(func(int) string)(100)
	if v != "str:100" {
		t.Failed()
	}
}
