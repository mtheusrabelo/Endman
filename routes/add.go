package routes

import "strings"

//AddRoute adds given route
func AddRoute(method string, path string) {
	path = strings.ToLower(path)
	method = strings.ToLower(method)

	switch method {
	case "get":
		r.GET(path, defaultGetHandler)
	case "post":
		r.POST(path, defaultHandler)
	case "put":
		r.PUT(path, defaultHandler)
	case "patch":
		r.PATCH(path, defaultHandler)
	case "delete":
		r.DELETE(path, defaultHandler)
	case "options":
		r.OPTIONS(path, defaultHandler)
	}
}
