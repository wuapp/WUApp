package wuapp

/*
#import "util.h"

void tlog(const char* msg) {
	printf("tlog before: %p,%s\n",msg,msg);
	goUILog("test:%s",msg);
	printf("tlog after: %p,%s\n",msg,msg);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func utilLog(msg string) {
	cMsg := C.CString(msg)
	C.free(unsafe.Pointer(cMsg))
	fmt.Printf("pointer %p\n", cMsg)
	C.tlog(cMsg)
}
