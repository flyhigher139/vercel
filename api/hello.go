package handler

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World!")
	w.Write([]byte("Hello World, again"))
}
