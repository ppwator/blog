package linkPack

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
	"unsafe"
)

func Test_memmove1(t *testing.T) {
	src := []byte{1, 2, 3, 4, 5, 6}
	dest := make([]byte, 10)

	spew.Dump(src)
	spew.Dump(dest)
	//// copy slice 1
	//srcp := (*GoSlice)(unsafe.Pointer(&src))
	//destp := (*GoSlice)(unsafe.Pointer(&dest))
	//
	//memmove(destp.Ptr, srcp.Ptr, unsafe.Sizeof(byte(0))*6)
	////copy slice 2
	copy(dest, src)

	spew.Dump(src)
	spew.Dump(dest)
}

func Test_memmove(t *testing.T) {
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
