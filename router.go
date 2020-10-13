package wuapp

import "regexp"

//type action = func(*Context)

type route struct {
	action func(*Context)
	paras  []string       //named parameters
	regex  *regexp.Regexp //if there are parameters
}

var (
	namedPart    = regexp.MustCompile(`:([^/]+)`)
	escapeRegExp = regexp.MustCompile(`([\-{}\[\]+?.,\\\^$|#\s])`)

	routes = make(map[string]*route)
)

func parseRoute(pattern string, route *route) {
	params := namedPart.FindAllStringSubmatch(pattern, -1)
	if params != nil {
		l := len(params)
		route.paras = make([]string, l)
		for i, param := range params {
			route.paras[i] = param[1]
		}
		pattern = namedPart.ReplaceAllString(pattern, `([^/]+)`)
		route.regex = regexp.MustCompile(pattern)
	}
	routes[pattern] = route
}

func dispatch(url string) (action func(*Context), params map[string]string) {
	for key, route := range routes {
		if route.regex == nil {
			if key == url {
				action = route.action
				return
			}

		} else {
			matches := route.regex.FindAllStringSubmatch(url, -1)
			if matches != nil && len(matches) == 1 {
				params = make(map[string]string)
				vals := matches[0][1:]
				l, lkey := len(vals), len(route.paras)
				if l > lkey {
					l = lkey
				}
				for i := 0; i < l; i++ {
					params[route.paras[i]] = vals[i]
				}
				action = route.action
				return
			}
		}
	}
	return
}
