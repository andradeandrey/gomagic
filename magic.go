package magic

// #include <magic.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type Magic_t C.magic_t

func Open(flags int) Magic_t {
	return (Magic_t)(C.magic_open(C.int(flags)))
}

func Close(cookie Magic_t) {
	C.magic_close((C.magic_t)(cookie))
}

func Error(cookie Magic_t) string {
	s := (C.magic_error((C.magic_t)(cookie)))
	return C.GoString(s)
}

func Errno(cookie Magic_t) int {
	return (int)(C.magic_errno((C.magic_t)(cookie)))
}

func File(cookie Magic_t, filename string) string {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return C.GoString(C.magic_file(cookie, cfilename))
}

func Buffer(cookie Magic_t, b []byte, length C.size_t) string {
	//return C.GoString(C.magic_buffer(cookie, C.void(*b), length))
	return "j"
}

func SetFlags(cookie Magic_t, flags int) int {
	return (int)(C.magic_setflags(cookie, C.int(flags)))
}

func Check(cookie Magic_t, filename string) int {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_check(cookie, cfilename))
}

func Compile(cookie Magic_t, filename string) int {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_compile(cookie, cfilename))
}

func Load(cookie Magic_t, filename string) int {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))
	return (int)(C.magic_compile(cookie, cfilename))
}
