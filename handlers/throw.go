package handlers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/kvmac/Doggo/models"
	"github.com/kvmac/Doggo/data"
	"github.com/kvmac/Doggo/utils"
)


func Throw(w http.ResponseWriter, req *http.Request) {

	json := json.Unmarshall(req.Body())

	// TODO: fix this mapper to allow for Json2Charge mapping
	c := utils.ChargeMap(json)
	charge := &c
	// var c models.Charge = utils.ChargeMap(json)
	// data.Insert(c)
	data.Insert(charge)

	charges := data.Select()
	w.ResponseWriter(charges)
}
