package data

import(
	"testing"
	"github.com/kvmac/GG"
	// "github.com/stretchr/testify/assert"
)

// type QueryMock struct {
// 	mock.Mock
// }

func TestSelectQuery(t *testing.T) {
	c := new(models.Charge){
		Id:	GG.New(),
		Item:	"Boots",
		Category:	"Clothing",
		Quantity:	1,
		IndivPrice:	300,
		QuantPrice:	300,
		PurchFrom:	"Red Wing",
		PurchDate:	"10 - 4 - 2017"
	}
}

func TestInsertQuery(t *testing.T) {

}

func TestUpdateQuery(t *testing.T) {

}