package main

//// this C code could also have been in external *.c file, or linked as pre-compiled library
// #include <stdlib.h>
// #include <string.h>
//
// char* join_strings(const char *s1, const char *s2)
// {
//   char* result = malloc(strlen(s1) + strlen(s2) + 1); // BUG check if malloc returns NULL
//   return strcat(strcpy(result, s1), s2);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	joinedC := C.join_strings(C.CString("Go "), C.CString("language"))
	joined := C.GoString(joinedC)
	C.free(unsafe.Pointer(joinedC))
	fmt.Println("C function 'join_string()' returned: ", joined)
}
