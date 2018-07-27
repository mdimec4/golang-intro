package main

//// this C code could also have been in external *.c file, or linked as pre-compiled library
// #include <stdlib.h>
// #include <string.h>
//
// char* join_strings(const char *s1, const char *s2)  {
//   char* joined = malloc(strlen(s1) + strlen(s2) + 1); // BUG check if malloc returns NULL
//   return strcat(strcpy(joined, s1), s2);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	c_s1, c_s2 := C.CString("Go "), C.CString("language")
	c_joined := C.join_strings(c_s1, c_s2)
	C.free(unsafe.Pointer(c_s1))
	C.free(unsafe.Pointer(c_s2))
	fmt.Println(C.GoString(c_joined))
	C.free(unsafe.Pointer(c_joined))
}
