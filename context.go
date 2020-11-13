package wua

import (
	"encoding/json"
	"net/url"
	"strconv"
)

const callback = "wuapp.callback"

type Context interface {
	GetParam(name string) string
	GetParamOr(name, defaultVal string) string
	GetUnescapedParam(name string) (val string)
	GetBoolParam(name string) (b bool, err error)
	GetBoolParamOr(name string, defaultVal bool) (b bool)
	GetIntParam(name string) (i int, err error)
	GetIntParamOr(name string, defaultVal int) (i int)
	GetFloatParam(name string) (f float64, err error)
	GetFloatParamOr(name string, defaultVal float64) (f float64)
	GetEntity(v interface{}) (err error)
	Success(feedback interface{})
	Error(feedback ...interface{})
}

type contextBase struct {
	//message *Message
	data   []byte
	params map[string]string
}

// GetParam get a string parameter from the url
func (ctx *contextBase) GetParam(name string) string {
	return ctx.params[name]
}

func (ctx *contextBase) GetParamOr(name, defaultVal string) string {
	val := ctx.params[name]
	if val == "" {
		return defaultVal
	}
	return val
}

// GetParam get a string parameter from the url
func (ctx *contextBase) GetUnescapedParam(name string) (val string) {
	val, err := url.PathUnescape(ctx.params[name])
	if err != nil {
		//Log("Unescape parameter", name, "failed:",err.Error())
	}
	return
}

// GetBoolParam get a bool parameter from the url
func (ctx *contextBase) GetBoolParam(name string) (b bool, err error) {
	str := ctx.GetParam(name)
	b, err = strconv.ParseBool(str)
	if err != nil {
		//Log("convert data to bool failed:", err)
	}
	return
}

// GetBoolParamOr get a bool parameter from the url
func (ctx *contextBase) GetBoolParamOr(name string, defaultVal bool) bool {
	str := ctx.GetParam(name)
	b, err := strconv.ParseBool(str)
	if err != nil {
		return defaultVal
	}
	return b
}

// GetIntParam get a int parameter from the url
func (ctx *contextBase) GetIntParam(name string) (i int, err error) {
	str := ctx.GetParam(name)
	i, err = strconv.Atoi(str)
	if err != nil {
		//Log("convert data to int failed:", err)
	}
	return
}

func (ctx *contextBase) GetIntParamOr(name string, defaultVal int) int {
	str := ctx.GetParam(name)
	i, err := strconv.Atoi(str)
	if err != nil {
		return defaultVal
	}
	return i

}

// GetFloatParam get a float parameter from the url
func (ctx *contextBase) GetFloatParam(name string) (f float64, err error) {
	str := ctx.GetParam(name)
	f, err = strconv.ParseFloat(str, 32)
	if err != nil {
		//Log("convert data to float failed:", err)
	}
	return
}

func (ctx *contextBase) GetFloatParamOr(name string, defaultVal float64) float64 {
	str := ctx.GetParam(name)
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		//Log("convert data to float failed:", err)
		return defaultVal
	}
	return f
}

// GetParam get an entity from the requested data
func (ctx *contextBase) GetEntity(v interface{}) (err error) {
	//err = json.Unmarshal([]byte(ctx.message.Data), v)
	err = json.Unmarshal(ctx.data, v)
	if err != nil {
		//Log("get entity failed:", err)
	}
	return
}
