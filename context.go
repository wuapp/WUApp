package wuapp

import (
	"encoding/json"
	"net/url"
	"strconv"
	"wuapp/util"
)

type Context struct {
	message *Message
	params          map[string]string
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

// feedback should be a primary type, or implement the fmt.Stringer interface
// if not, convert your value to string first. e.g. string(bytes)
func (ctx *Context) Success(feedback ...interface{}) {
	//if feedback.(type) == byte {}
	if ctx.message.Success != "" {
		invokeJavascript(formJsCallString(ctx.message.Success,feedback))
	}
}

func (ctx *Context) Error(err ...interface{}) {
	if ctx.message.Error != "" {
		invokeJavascript(formJsCallString(ctx.message.Error,err))
	}
}

func formJsCallString(funcName string, args []interface{}) string  {
	return funcName + util.Join(args,",","(",")","'")
}

