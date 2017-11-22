package models

import(
	"github.com/kvmac/GG"
)

type Charge struct {
	Id	GG.Guid
	Item	string
	Category	string
	Quantity	int
	IndivPrice	int
	QuantPrice	int
	PurchFrom	string
	PurchDate	string
}
