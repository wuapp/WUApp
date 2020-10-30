//+build !web

package wuapp

/*
#include <stdlib.h>
#include "window.h"

*/
import "C"
import (
	"encoding/json"
	"github.com/wuapp/util"
	"os"
	"path"
	"runtime"
	"unsafe"
)

const defaultDir = "ui"
const defaultIndex = "index.html"
const clientHandler = "wuapp.receive"

var menuDefs []MenuDef = nil
var windowSettings = WindowSettings{UIDir: defaultDir,
	Index: defaultIndex,
}

func BoolToCInt(b bool) (i C.int) {
	if b {
		i = 1
	}
	return
}

func AddMenu(menuDefArray []MenuDef) {
	menuDefs = menuDefArray
}

func convertSettings(settings WindowSettings) C.WindowSettings {
	//dir := path.Dir(settings.Url)
	if settings.UIDir == "" {
		settings.UIDir = defaultDir
	}

	if settings.Index == "" {
		settings.Index = defaultIndex
	}

	if settings.Url == "" {
		settings.Url = path.Join(settings.UIDir, settings.Index)
		if runtime.GOOS == "linux" {
			wd, _ := os.Getwd()
			settings.Url = path.Join("file://", wd, settings.Url)
		} else if runtime.GOOS == "android" {
			settings.Url = path.Join("file:///android_asset/", settings.Url)
		}
	}

	// windows needs WebDir and Index
	// macOS and iOS need Url

	return C.WindowSettings{C.CString(settings.Title),
		C.CString(settings.UIDir),
		//C.CString(abs),
		C.CString(settings.Index),
		C.CString(settings.Url),
		C.int(settings.Left),
		C.int(settings.Top),
		C.int(settings.Width),
		C.int(settings.Height),
		BoolToCInt(settings.Resizable),
		BoolToCInt(settings.Debug),
	}
}

func create(settings WindowSettings) error {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := convertSettings(settings)
	cMenuDefs, count := convertMenuDefs(menuDefs)
	cCreate(cs, cMenuDefs, count)
	return nil
}

func activate() {

}

func sendMessage(msg string) {
	send(clientHandler, msg)
}

func send(funcName string, args ...interface{}) {
	for i, a := range args {
		Logger.Info("i:", i, "a:", a)
	}
	js := funcName + util.JoinEx(args, "(", ",", ")", `"`)
	cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	Logger.Info("invokeJavascript:", js)
	cInvokeJS(cJs, 0)
}

func invokeJavascriptAsync(js string) {
	cJs := C.CString(js)
	defer C.free(unsafe.Pointer(cJs))
	cInvokeJS(cJs, 1)
}

//export receive
func receive(msg *C.char) {
	goMsg := C.GoString(msg)
	//defer C.free(unsafe.Pointer(msg))
	Logger.Info("ClientHandler:", goMsg)
	message := new(Message)
	err := json.Unmarshal([]byte(goMsg), message)
	if err != nil {
		//Log("unmarshal error:", err)
		return
	}

	ctx := newContext(message)
	dispatch(message.Url, ctx)
}

func exit() {
	cExit()
}
