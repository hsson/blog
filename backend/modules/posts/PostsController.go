package posts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Machiel/slugify"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
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

func New(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	if u == nil || !u.Admin {
		log.Warningf(ctx, "Unauthorized API attempt")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var p Post
	err := decoder.Decode(&p)
	if err != nil {
		log.Errorf(ctx, "Could not decode new post")
		http.Error(w, "Couldn't decode new post", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if p.Title == "" || p.Body == "" {
		http.Error(w, "No title or body specified", http.StatusBadRequest)
		return
	}
	var nullTime time.Time
	if p.PostDate == nullTime {
		p.PostDate = time.Now()
	}
	p.Slug = slugify.Slugify(p.Title)

	// Input is okay, now save
	err = p.save(ctx)
	if err != nil {
		log.Errorf(ctx, "Could not save post to database")
		http.Error(w, "Could not save post to database", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "200 OK")
}

func Edit(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	if u == nil || !u.Admin {
		log.Warningf(ctx, "Unauthorized API attempt")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	postSlug := mux.Vars(r)["slug"]
	var oldPost Post

	err := datastore.Get(ctx, keyToString(ctx, postSlug), &oldPost)
	if err != nil {
		log.Infof(ctx, "Couldn't find post in database")
		http.Error(w, "Post does not exist", http.StatusBadRequest)
		return
	}

	var newPost Post
	err = json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		log.Errorf(ctx, "Could not decode post")
		http.Error(w, "Couldn't decode post", http.StatusBadRequest)
		return
	}
	oldPost.Slug = postSlug
	if newPost.Title != "" {
		oldPost.Title = newPost.Title
	}
	if newPost.Body != "" {
		oldPost.Body = newPost.Body
	}
	err = oldPost.save(ctx)
	if err != nil {
		log.Errorf(ctx, "Could not save post to database")
		http.Error(w, "Could not save post to database", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "200 OK")
}
