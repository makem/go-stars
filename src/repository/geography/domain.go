package geography

import (
	. "domain"
)

//implementation
type country struct {
	CodeΩ string `bson:"_id" json:"code"`
	NameΩ string `bson:"name" json:"name"`
}

func newCountry(code, name string) ICountry {
	return &country{
		CodeΩ: code,
		NameΩ: name,
	}
}

func (this country) Id() string {
	return this.CodeΩ
}

func (this *country) Name(Ω ...string) string {
	if len(Ω) > 0 {
		this.NameΩ = Ω[0]
	}
	return this.NameΩ
}
