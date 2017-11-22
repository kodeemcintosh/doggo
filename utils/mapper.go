package utils

import(
	"github.com/kvmac/GG"
	"github.com/kvmac/Doggo/models"
)

func ChargeMap(id GG.Guid, item, purchased_from, purchased_date string, quantity, individual_price, quantity_price int) (models.Charge) {

	// c := &Charge{
	c := &models.Charge{
		Id:	id,
		Item:	item,
		Quantity:	quantity,
		IndivPrice:	individual_price,
		QuantPrice:	quantity_price,
		PurchFrom:	purchased_from,
		PurchDate:	purchased_date,
	}
	// return *c
	return *c
}