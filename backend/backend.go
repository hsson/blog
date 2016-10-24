package backend

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type Post struct {
  Id        int     `json:"id"`
  PostDate  int     `json:"postDate"`
  Title     string  `json:"title"`
  Body      string  `json:"body"`
}

var posts = []Post {
  {1, 1477336203000, "Hello World! 2", "I am a body"},
  {2, 1477336203000, "Hello World! 2", "I am a body"},
}

func init() {
  http.HandleFunc("/api/posts", getPosts)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
  enc := json.NewEncoder(w)
  err := enc.Encode(posts)
  if (err != nil) {
    fmt.Fprintln(w, "Something went wrong")
  }
}
