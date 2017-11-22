package data

import(
	"github.com/kvmac/Doggo/models"
	_ "github.com/lib/pq"
)

func UpdateChargeById(c models.Charge) {
	a.Initialize()

	if c.Id == nil {
		c.Id.New()
	}

	sql := `UPDATE Charges SET item = $2,
				quantity = $3,
				individual_price = $4,
				quantity_price = $5,
				purchased_from = $6,
				purchased_date = $7,
				last_modified = $8
				WHERE id = $1;`

	_, err := a.DB.Exec(sql, c.Id, c.Item, c.Quantity, c.IndivPrice, c.QuantPrice, c.PurchFrom, c.PurchDate, time.Now())
	utils.ErrCheck(err)

	a.Finalize()
}
