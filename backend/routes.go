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
	Route{"IndexPosts", "GET", "/api/posts", false, posts.Index},
	Route{"GetPost", "GET", "/api/posts/{id:[0-9]+}", false, posts.Get},
	//Route{"PopulateDummyData", "GET", "/api/posts/dummy", false, posts.PopulateDummyData},
}
