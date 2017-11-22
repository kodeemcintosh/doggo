package data

import(
	"time"
	// "github.com/kvmac/GG"
	"github.com/kvmac/Doggo/models"
	_ "github.com/lib/pq"
)

func InsertNewCharge(c *models.Charge) {

	a := models.App{}

	a.Initialize()

	if c.Id == nil {
		c.Id.New()
	}


	sql := `INSERT INTO Charges (id,
				 item,
				 category,
				 quantity, 
				 individual_price, 
				 quantity_price, 
				 purchased_from, 
				 last_modified) 
				 VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := a.DB.Exec(sql, c.Id, c.Item, c.Category, c.Quantity, c.IndivPrice, c.QuantPrice, c.PurchFrom, c.PurchDate, time.Now())
	utils.ErrCheck(err)

	a.Finalize()
}