// +build !android

package wuapp

/*
#include <stdlib.h>
*/
import "C"
import (
	"github.com/wuapp/log"
)

var Logger log.Logger

func init() {
	Logger = log.GetLogger()
}

//export goLog
func goLog(msg *C.char) {
	s := C.GoString(msg)
	Logger.Info(s)
}

//export goLogError
func goLogError(msg *C.char) {
	s := C.GoString(msg)
	Logger.Error(s)
}
