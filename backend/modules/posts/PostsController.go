package posts

import (
  "net/http"
  "encoding/json"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
  "time"
)

var dummyPosts = []Post {
  {1, time.Date(2016, time.October, 24, 18, 31, 01, 0, time.Local), "Hello World!", "I am a the first post"},
  {2, time.Date(2015, time.June, 3, 15, 21, 07, 0, time.Local), "Hello World! 2", "I am a body"},
}

func PopulateDummyData(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  for _,dummyPost := range dummyPosts {
    dummyPost.save(ctx)
  }
}

func Index(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  posts,err := getPosts(ctx)

  if err != nil {
    log.Errorf(ctx, "Couldn't fetch blog posts from database")
    http.Error(w, "Couldn't fetch blog posts", http.StatusInternalServerError)
  }

  enc := json.NewEncoder(w)
  err = enc.Encode(posts)
  if (err != nil) {
    log.Errorf(ctx, "Couldn't encode posts to json")
    http.Error(w, "Couldn't encode posts to json", http.StatusInternalServerError)
  }
}
