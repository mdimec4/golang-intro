package main

//// this C code could also have been in external *.c file,
//// or linked as pre-compiled library
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
//
// void hello(void)
// {
//   printf("hello from C\n");
// }
//
// char* join_strings(const char *s1, const char *s2)
// {
//   char* result;
//
//   result = malloc(strlen(s1) + strlen(s2) + 1);
//   if(result == NULL)
//     return NULL;
//   strcpy(result, s1);
//   strcat(result, s2);
//   return result;
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	C.hello()
	// it seams that Go presentation code runner "present",
	// prints nothing from <stdio.h> without `fflush()`
	C.fflush(C.stdout)

	joinedC := C.join_strings(C.CString("Go "), C.CString("language"))
	joined := C.GoString(joinedC)
	C.free(unsafe.Pointer(joinedC))
	fmt.Println("C function 'join_string()' returned: ", joined)
}
