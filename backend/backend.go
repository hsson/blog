package backend

import (
	"net/http"
)

func Run() {
	r := NewRouter()
	http.Handle("/", r)
}
