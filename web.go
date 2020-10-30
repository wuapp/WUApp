// +build web

package wuapp

import (
	"net/http"
	"path"
	"strings"
)

func AddMenu(menuDefArray []MenuDef) {

}

type Service struct {
}

func (svc *Service) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.RequestURI == "*" {
		if req.ProtoAtLeast(1, 1) {
			rw.Header().Set("Connection", "close")
		}
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	//originControl(rw) todo
	if req.Method == "OPTIONS" {
		return
	}

	url := req.URL.Path
	if strings.HasPrefix(url, "/ws") {
		ctx := newContext(rw, req)
		dispatch(url[4:], ctx)
	} else {
		if !strings.HasPrefix(url, "/") {
			url = "/" + url
			req.URL.Path = url
		}
		if url == "/" {
			url = "index.html"
		}
		url := path.Join("./ui", url)
		http.ServeFile(rw, req, url)
	}
}

func create(settings WindowSettings) error {
	port := Config.GetStringOr("port", ":8088")
	//go http.ListenAndServe(":8888",http.FileServer(http.Dir("./ui")))
	return http.ListenAndServe(port, new(Service))
}

func exit() {

}
