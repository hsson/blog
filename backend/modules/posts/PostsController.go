package posts

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

var dummyPosts = []Post{
	{"hello-world", time.Date(2016, time.October, 24, 18, 31, 01, 0, time.Local), "Hello World!", "I am a the first post"},
	{"hello-world-2", time.Date(2015, time.June, 3, 15, 21, 07, 0, time.Local), "Hello World! 2", "I am a body"},
}

func Generate(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	for _, dummyPost := range dummyPosts {
		dummyPost.save(ctx)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	posts, err := getPosts(ctx)

	if err != nil {
		log.Errorf(ctx, "Couldn't fetch blog posts from database")
		http.Error(w, "Couldn't fetch blog posts", http.StatusInternalServerError)
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(posts)
	if err != nil {
		log.Errorf(ctx, "Couldn't encode posts to json")
		http.Error(w, "Couldn't encode posts to json", http.StatusInternalServerError)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postSlug := vars["slug"]
	ctx := appengine.NewContext(r)
	post, err := getPost(ctx, postSlug)
	if err != nil {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	err = enc.Encode(post)
	if err != nil {
		log.Errorf(ctx, "Couldn't encode post to json")
		http.Error(w, "Couldn't encode post to json", http.StatusInternalServerError)
	}
}
