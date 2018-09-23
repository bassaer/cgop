package main

/*
#include <stdlib.h>
extern void hello(const char*);
*/
import "C"
import "unsafe"

func main() {
	name := C.CString("foo")
	defer C.free(unsafe.Pointer(name))
	C.hello(name)
}
