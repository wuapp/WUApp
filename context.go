package wuapp

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const callback = "wuapp.callback"

type Context struct {
	message *Message
	params  map[string]string
}

// GetParam get a string parameter from the url
func (ctx *Context) GetParam(name string) string {
	return ctx.params[name]
}

// GetParam get a string parameter from the url
func (ctx *Context) GetUnescapedParam(name string) (val string) {
	val, err := url.PathUnescape(ctx.params[name])
	if err != nil {
		//Log("Unescape parameter", name, "failed:",err.Error())
	}
	return
}

// GetBoolParam get a bool parameter from the url
func (ctx *Context) GetBoolParam(name string) (b bool, err error) {
	str := ctx.GetParam(name)
	b, err = strconv.ParseBool(str)
	if err != nil {
		//Log("convert data to bool failed:", err)
	}
	return
}

// GetIntParam get a int parameter from the url
func (ctx *Context) GetIntParam(name string) (i int, err error) {
	str := ctx.GetParam(name)
	i, err = strconv.Atoi(str)
	if err != nil {
		//Log("convert data to int failed:", err)
	}
	return
}

func (ctx *Context) GetIntParamOr(name string, defaultVal int) (i int) {
	str := ctx.GetParam(name)
	var err error
	i, err = strconv.Atoi(str)
	if err != nil {
		i = defaultVal
	}
	return

}

// GetFloatParam get a float parameter from the url
func (ctx *Context) GetFloatParam(name string) (f float64, err error) {
	str := ctx.GetParam(name)
	f, err = strconv.ParseFloat(str, 32)
	if err != nil {
		//Log("convert data to float failed:", err)
	}
	return
}

// GetParam get an entity from the requested data
func (ctx *Context) GetEntity(v interface{}) (err error) {
	err = json.Unmarshal([]byte(ctx.message.Data), v)
	if err != nil {
		//Log("get entity failed:", err)
	}
	return
}

func (ctx *Context) Done(ok bool, feedback interface{}) {
	invokeJavascript(callback, ctx.message.Id, ok, feedback)
}

// feedback should be a primary type, or implement the fmt.Stringer interface
// if not, convert your value to string first. e.g. string(bytes)
func (ctx *Context) Success(feedback interface{}) {
	Log("success:", feedback)
	ctx.Done(true, feedback)
}

func (ctx *Context) Error(feedback ...interface{}) {
	ctx.Done(false, feedback)
}
