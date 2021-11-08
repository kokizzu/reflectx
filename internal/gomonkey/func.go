package gomonkey

import (
	"reflect"
	"syscall"
	"unsafe"
)

func MakeFunc(i interface{}, typ reflect.Type, fn func([]reflect.Value) []reflect.Value) reflect.Value {
	target := reflect.ValueOf(i)
	double := reflect.MakeFunc(typ, fn)
	replace(*(*uintptr)(getPointer(target)), uintptr(getPointer(double)))
	p := (*funcValue)(unsafe.Pointer(&target))
	r := (*funcValue)(unsafe.Pointer(&double))
	r.p = p.p
	return *(*reflect.Value)(unsafe.Pointer(r))
}

func replace(target, double uintptr) []byte {
	code := buildJmpDirective(double)
	bytes := entryAddress(target, len(code))
	original := make([]byte, len(bytes))
	copy(original, bytes)
	modifyBinary(target, code)
	return original
}

type funcValue struct {
	_ uintptr
	p unsafe.Pointer
}

func getPointer(v reflect.Value) unsafe.Pointer {
	return (*funcValue)(unsafe.Pointer(&v)).p
}

func entryAddress(p uintptr, l int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{Data: p, Len: l, Cap: l}))
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
