package wua

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

const (
	//errInternalError = "500 internal server error"
	//errNotFount = "404 page not found"
	//errMethodNotAllowed = "405 method not allowed"
	contentType = "Content-Type"
	typeJson    = "application/json; charset=utf-8"
	typeHtml    = "text/html; charset=utf-8"
	typeXML     = "text/xml; charset=utf-8"
)

type HttpContext struct {
	*contextBase
	req *http.Request
	rw  http.ResponseWriter
}

func newHttpContext(rw http.ResponseWriter, req *http.Request) Context {
	body, _ := ioutil.ReadAll(req.Body)
	return &HttpContext{&contextBase{data: body}, req, rw}
}

type Feedback struct {
	OK       bool        `json:"ok"`
	Feedback interface{} `json:"feedback"`
}

func (ctx *HttpContext) Done(ok bool, feedback interface{}) {
	f := Feedback{ok, feedback}
	ctx.ServeJson(f)
}

// feedback should be a primary type, or implement the fmt.Stringer interface
// if not, convert your value to string first. e.g. string(bytes)
func (ctx *HttpContext) Success(feedback interface{}) {
	Logger.Info("success:", feedback)
	ctx.Done(true, feedback)
}

func (ctx *HttpContext) Error(feedback ...interface{}) {
	ctx.Done(false, feedback)
}

func (ctx *HttpContext) GetPlainReq() (body []byte) {
	body, err := ioutil.ReadAll(ctx.req.Body)
	if err != nil {
		return
	}
	return
}

func (ctx *HttpContext) GetReqHeader(key string) (val string) {
	return ctx.req.Header.Get(key)
}

func (ctx *HttpContext) SetHeader(key, val string) {
	ctx.rw.Header().Set(key, val)
}

func (ctx *HttpContext) ServeBody(content []byte) {
	//todo:gzip
	ctx.rw.Write(content)
}

func (ctx *HttpContext) ServeString(o ...interface{}) {
	//fmt.Fprint(ctx.rw, o)
	s := fmt.Sprint(o...)
	ctx.rw.Write([]byte(s))
}

func (ctx *HttpContext) ServeJson(o interface{}) {
	bs, err := json.Marshal(o)
	if err != nil {
		ctx.InternalError()
	}
	ctx.SetHeader(contentType, typeJson)
	Logger.Info("ServeJson:", string(bs))
	ctx.rw.Write(bs)
}

func (ctx *HttpContext) ServeHtml(content []byte) {
	ctx.SetHeader(contentType, typeHtml)
	ctx.rw.Write(content)
}

func (ctx *HttpContext) ServeXML(content []byte) {
	ctx.SetHeader(contentType, typeXML)
	ctx.rw.Write(content)
}

func (ctx *HttpContext) ServeStatic(name string, content []byte) {
	ctype := mime.TypeByExtension(filepath.Ext(name))
	if ctype == "" {
		ctype = typeHtml
	}
	ctx.SetHeader(contentType, ctype)
	ctx.rw.Write(content)
}

func (ctx *HttpContext) ServeHtmlFile(path string) {
	ctx.writeFile(path, typeHtml)
}

func (ctx *HttpContext) ServeFile(path string) {
	ctype := mime.TypeByExtension(filepath.Ext(path))
	if ctype == "" {
		ctype = typeHtml
	}
	ctx.writeFile(path, ctype)
}

func (ctx *HttpContext) writeFile(path, fileType string) {
	ctx.SetHeader(contentType, fileType)
	f, err := os.Open(path)
	if err != nil {
		ctx.ServeStatus(http.StatusNotFound)
		return
	}
	io.Copy(ctx.rw, f)
}

func (ctx *HttpContext) ServeByTemplate(templ *template.Template, data interface{}) {
	ctx.SetHeader(contentType, typeHtml)

	var err error
	/*if compress {
		ctx.SetHeader("Content-Encoding", "gzip")
		w := gzip.NewWriter(ctx.rw)
		err = templ.Execute(w, data) // templates.ExecuteTemplate(w, templ, data)
		defer w.Close()
	} else {*/
	//err = templates.ExecuteTemplate(ctx.rw, templ, data)
	err = templ.Execute(ctx.rw, data)
	//}
	//log.Println("templ:",templ,"data:",data)

	if err != nil {
		//http.Error(ctx.rw, err.Error(), http.StatusInternalServerError)
		ctx.InternalError()
		//getLogger().Error(ctx.req, "Error sending response", err)
	}
}

func (ctx *HttpContext) ServeStatus(code int) {
	ctx.rw.WriteHeader(code)
}

func (ctx *HttpContext) Ok() {
	ctx.ServeStatus(http.StatusOK)
}

func (ctx *HttpContext) InternalError() {
	ctx.ServeStatus(http.StatusInternalServerError)
}

func (ctx *HttpContext) OkOrError(ok bool) {
	if ok {
		ctx.Ok()
	} else {
		ctx.InternalError()
	}
}

func (ctx *HttpContext) BadRequest() {
	ctx.ServeStatus(http.StatusBadRequest)
}

func (ctx *HttpContext) Unauthorized() {
	ctx.ServeStatus(http.StatusUnauthorized)
}

func (ctx *HttpContext) PageNotFound() {
	//if no custom handler
	ctx.ServeStatus(http.StatusNotFound)
}

func (ctx *HttpContext) Conflict() {
	ctx.ServeStatus(http.StatusConflict)
}

func (ctx *HttpContext) MethodNotAllowed() {
	ctx.ServeStatus(http.StatusMethodNotAllowed)
}
