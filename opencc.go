package opencc

// #cgo pkg-config: opencc
// #include <stdlib.h>
// #include "opencc/opencc.h"
import "C"

import (
	"errors"
	"math"
	"unsafe"
)

type Converter struct {
	ptr C.opencc_t
}

func NewConverter(configFileName string) (*Converter, error) {
	cfg := C.CString(configFileName)
	defer C.free(unsafe.Pointer(cfg))
	ptr := C.opencc_open(cfg)

	if uintptr(unsafe.Pointer(ptr)) == 0xffffffffffffffff {
		msg := C.GoString(C.opencc_error())
		return nil, errors.New(msg)
	}

	return &Converter{ptr: ptr}, nil
}

/*
func (c *Converter) Convert(input string) string {
    cInput := C.CString(input)
    defer C.free(unsafe.Pointer(cInput))

    x := C.opencc_convert_utf8(c.ptr, cInput, C.size_t(len(input)))
    defer C.opencc_convert_utf8_free(x)

    return C.GoString(x)
}

func (c *Converter) ConvertBytes(input []byte) {
    x := C.opencc_convert_utf8(c.ptr, (*C.char)(unsafe.Pointer(&input[0])), C.size_t(len(input)))
    defer C.opencc_convert_utf8_free(x)

    y := (*C.char)(unsafe.Pointer(&x))

    log.Print(string(y))
}
*/

func (c *Converter) Convert(input []byte) []byte {
	// https://karthikkaranth.me/blog/calling-c-code-from-go/
	ptr := C.malloc(C.sizeof_char * C.ulong(len(input)+int(math.Ceil(float64(len(input))/1024))*256))
	defer C.free(unsafe.Pointer(ptr))

	cInput := C.CBytes(input)
	defer C.free(cInput)

	size := C.opencc_convert_utf8_to_buffer(c.ptr,
		(*C.char)(cInput), C.size_t(len(input)), (*C.char)(ptr))

	return C.GoBytes(ptr, C.int(size))
}

func (c *Converter) Close() {
	C.opencc_close(c.ptr)
}
