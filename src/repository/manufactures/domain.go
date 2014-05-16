package manufactures

import (
	. "domain"
	"time"
)

type manufacture struct {
	NameΩ    string    `bson:"_id"`
	ActiveΩ  bool      `bson:"active"`
	AddressΩ IAddress  `bson:"address"`
	OpenedΩ  time.Time `bson:"opened"`
}

func NewManufacture(name string, address IAddress) IManufacture {
	return &manufacture{
		NameΩ:    name,
		AddressΩ: address,
		OpenedΩ:  time.Now(),
		ActiveΩ:  false,
	}
}

func (manufacture *manufacture) Name(name ...string) string {
	if len(name) > 0 {
		manufacture.NameΩ = name[0]
	}
	return manufacture.NameΩ
}

func (manufacture *manufacture) Active(active ...bool) bool {
	if len(active) > 0 {
		manufacture.ActiveΩ = active[0]
	}
	return manufacture.ActiveΩ
}

func (manufacture *manufacture) Address(address ...IAddress) IAddress {
	if len(address) > 0 {
		manufacture.AddressΩ = address[0]
	}
	return manufacture.AddressΩ
}

func (manufacture *manufacture) Opened() time.Time {
	return manufacture.OpenedΩ
}
