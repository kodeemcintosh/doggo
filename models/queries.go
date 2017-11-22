package models

type queries interface {
	Select() (Charges)
	Insert(*Charge) (Charges)
	Update(*Charge) (Charges)
}