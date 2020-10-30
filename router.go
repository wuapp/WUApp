package wuapp

import "regexp"

var (
	namedPart = regexp.MustCompile(`:([^/]+)`)
	//discard \*.*
	//custom regexp.MustCompile(`\$(.+)<(.+)>`)
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

func dispatch(url string, ctx *Context) {
	var action action
	for key, route := range routes {
		if route.regex == nil {
			if key == url {
				action = route.action
				goto A
			}
		} else {
			matches := route.regex.FindAllStringSubmatch(url, -1)
			if matches != nil && len(matches) == 1 {
				vals := matches[0][1:]
				l, lkey := len(vals), len(route.paras)
				if l > lkey {
					l = lkey
				}
				for i := 0; i < l; i++ {
					ctx.params[route.paras[i]] = vals[i]
				}
				action = route.action
				goto A
			}
		}
	}
A:
	if action != nil {
		action(ctx)
	} else {
		ctx.Error("Function not found ")
	}

}
