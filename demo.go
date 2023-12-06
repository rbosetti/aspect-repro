package demo

// #include "unsafe.h"
import "C"

import (
	"unsafe"
)

func UseUnsafeInt() int {
	return int(uintptr(unsafe.Pointer(C.UseUnsafeInt))) + int(C.UseUnsafeInt())
}
