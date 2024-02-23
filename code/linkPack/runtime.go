package main

import (
	"github.com/davecgh/go-spew/spew"
	"unsafe"
)

// //goland:noinspection GoUnusedParameter
//
//go:noescape
//go:linkname memmove runtime.memmove
func memmove(to unsafe.Pointer, from unsafe.Pointer, n uintptr)

type GoSlice struct {
	Ptr unsafe.Pointer
	Len int
	Cap int
}

type GoString struct {
	Ptr unsafe.Pointer
	Len int
}

func main() {
	str := "pedro"
	// 注意：这里的len不能为0，否则数据没有分配，就无法复制
	data := make([]byte, 10)
	spew.Dump(str)
	spew.Dump(data)

	memmove((*GoSlice)(unsafe.Pointer(&data)).Ptr, (*GoString)(unsafe.Pointer(&str)).Ptr,
		unsafe.Sizeof(byte(0))*5)
	spew.Dump(str)
	spew.Dump(data)
}
