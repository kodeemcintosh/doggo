package main

import (
	// "database/sql"
	// "flag"
	// "fmt"
	// "os"
	"github.com/gorilla/mux"
	"github.com/kvmac/Doggo/server"
	"encoding/json"
	"time"
	"net/http"
)

func main() {
	server.ServeDoggo()

	// r := mux.NewRouter()
	// r.handle("/throw", ThrowBall)
	// r.handle("/return", ReturnBall)
	// http.ListenAndServe(r, ":8080")

}
