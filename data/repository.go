package data

import(
	"github.com/kvmac/Doggo/models"
	"github.com/kvmac/Doggo/utils"
)


func Select() (models.Charges) {
	// Send the charge list client side
	charges := queries.SelectAllCharges()
	return charges
}

func Insert(c *models.Charge) (models.Charges) {
	// utils.ErrCheck(c)
	queries.InsertNewCharge(&c)

	// Refresh the charge list on the client side
	charges := queries.SelectAllCharges()
	return charges
}

func Update(c *models.Charge) (models.Charges) {
	// utils.ErrCheck(c)
	queries.UpdateChargeById(&c)

	// Refresh the charge list on the client side
	charges := queries.SelectAllCharges()
	return charges
}
