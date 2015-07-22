package system

// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

func System(cmd string) int {
	ptr := C.CString(cmd)
	defer C.free(unsafe.Pointer(ptr))
	return int(C.system(ptr))
}
