package handlers

import(
	"encoding/json"
	"net/http"
	"github.com/kvmac/Doggo/data"
	"github.com/kvmac/Doggo/models"
	"github.com/kvmac/Doggo/utils"
)


func Retrieve(w http.ResponseWriter, req *http.Request) {


	json := json.Unmarshall(req.Body())
	// var c models.Charge = utils.ChargeMap(json)
	c := utils.ChargeMap(json)
	charge := &c
	
	// Sends a pointer to the 'c' variable
	// data.Update(&c)
	data.Update(charge)

	charges := data.Select()

	w.ResponseWriter(charges)
}
