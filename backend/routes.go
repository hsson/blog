package backend

import (
	"net/http"

	"github.com/hsson/blog/backend/modules/posts"
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
	Route{"PopulateDummyData", "GET", "/api/posts/generate", false, posts.Generate},
	Route{"GetPost", "GET", "/api/posts/{slug}", false, posts.Get},
	Route{"NewPost", "POST", "/api/admin/posts", false, posts.New},
	Route{"EditPost", "PUT", "/api/admin/posts/{slug}", false, posts.Edit},
}
