package server

import(
	"github.com/gorilla/mux"
	"github.com/kvmac/Doggo/handlers"
	"net/http"
)

func ServeDoggo() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/throw", handlers.Throw)
	r.HandleFunc("/retrieve", handlers.Retrieve)
	r.HandleFunc("/fetch", handlers.Fetch)

	http.ListenAndServe(r, ":8080")
}