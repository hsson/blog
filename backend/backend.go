package backend

import (
  "net/http"
)

func init() {
  r := NewRouter()
  http.Handle("/", r)
}
