package backend

import (
	"net/http"
	"modules/posts"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Secured     bool
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Posts
	Route{"GetPosts", "GET", "/api/posts", false, posts.Index},
}
