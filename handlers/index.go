package handlers

import(
	"fmt"
	"net/http"
	"github.com/kvmac/Doggo/data"
)

func Index(w http.WriteResponse, req *http.Request) {

	charges := data.Select()


	w.WriteResponse(charges)

	fmt.Println("Index works")

}