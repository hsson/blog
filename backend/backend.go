package backend

import (
  "fmt"
  "net/http"
)

func init() {
  http.HandleFunc("/api/", getIndex)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello world!")
}
