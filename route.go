package wuapp

import "regexp"

type Encoding int

const (
	None Encoding = iota
	Gzip
	Deflate
)

type HttpMethod string

const (
	NotSupported HttpMethod = "" //not http
	GET          HttpMethod = "GET"
	PUT          HttpMethod = "PUT"
	POST         HttpMethod = "POST"
	DELETE       HttpMethod = "DELETE"
	HEAD                    = "HEAD"
)

type action func(*Context)

type route struct {
	url string
	//	encoding Encoding
	httpMethod HttpMethod
	action     action
	paras      []string       //named parameters
	regex      *regexp.Regexp //if there are parameters
}

func (r *route) addPara(para string) {
	if r.paras == nil {
		r.paras = make([]string, 1)
		r.paras[0] = para
	} else {
		r.paras = append(r.paras, para)
	}
}

/*
func Rpc(name string, execFunc func()) *route {
	return &route{name,Gob,None,execFunc}
}*/

func Route(url string, action action) {
	parseRoute(url, &route{url: url, httpMethod: GET, action: action})
}

func Get(url string, action action) {
	Route(url, action)
}

func PostRoute(url string, action action) {
	parseRoute(url, &route{url: url, httpMethod: POST, action: action})
}

func PutRoute(url string, action action) {
	parseRoute(url, &route{url: url, httpMethod: PUT, action: action})
}

func DeleteRoute(url string, action action) {
	parseRoute(url, &route{url: url, httpMethod: DELETE, action: action})
}
