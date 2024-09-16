package models

type Trip struct {
	Rider IRider
	Cab   ICab
	Price float64
	From  Location
	To    Location
}
