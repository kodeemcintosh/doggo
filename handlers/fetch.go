package handlers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/kvmac/Doggo/models"
	"github.com/kvmac/Doggo/data"
)

// GET request
func Fetch(w http.ResponseWriter, req *http.Request) {

	charges := data.Select()

	fmt.Println(req.Body)
	fmt.Println("fetch also works")
	w.Write(charges)

}