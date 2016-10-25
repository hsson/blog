package posts

import (
  "net/http"
  "encoding/json"
  "google.golang.org/appengine"
  "google.golang.org/appengine/log"
)

var posts = []Post {
  {1, 1477336203000, "Hello World!", "I am a the first post"},
  {2, 1477336203001, "Hello World! 2", "I am a body"},
}

func Index(w http.ResponseWriter, r *http.Request) {
  ctx := appengine.NewContext(r)
  enc := json.NewEncoder(w)
  err := enc.Encode(posts)
  if (err != nil) {
    log.Errorf(ctx, "Couldn't encode posts to json")
  }
}
