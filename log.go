// +build !android

package wuapp

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"io"
	"os"
	"unsafe"
)

//export goLog
func goLog(msg *C.char) {
	//s := fmt.Sprintln(msg)
	s := C.GoString(msg)
	defer C.free(unsafe.Pointer(msg))
	//`fmt.Printf("go log: %s\n",s)

	doLog(s)
}

func Log(args ...interface{}) {
	s := fmt.Sprintln(args...)
	doLog(s)
}

func Logf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	doLog(s)
}

var logger io.WriteCloser

func doLog(msg string) {
	if logger == nil {
		filename := Config.GetString("logFile")
		if filename != "" {
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0644))
			if err == nil {
				logger = file
			} else {
				logger = os.Stdout
			}
		} else {
			logger = os.Stdout
		}
	}
	go logger.Write([]byte(msg))
}

func closeLogger() {
	if logger != nil {
		logger.Close()
	}
}
