package data

import(
	"github.com/kvmac/Doggo/models"
	"github.com/kvmac/GG"
	_ "github.com/lib/pq"
)

func SelectAllCharges() (models.Charges) {
	a := models.App{}

	a.Initialize()

	sql := `SELECT * FROM Charges;`
	rows := a.DB.Query(sql)
	defer rows.Close()

	for rows.Next() {

		// var id, item, category, purchased_from, purchased_date string
		var item, category, purchased_from, purchased_date string
		var quantity, individual_price, quantity_price int
		var id GG.Guid

		err = rows.Scan(&id, &item, &category, &quantity, &individual_price, &quantity_price, &purchased_from, &purchased_date)

		c := utils.ChargeMap(id, item, category, quantity, individual_price, quantity_price, purchased_from, purchased_date)

		charges := &Charges{}
		// a.Charges = append(c, len+1)

		length := len(charges)

		charges = append(c, length)

		a.charges = charges
	}

	a.Finalize()
}
