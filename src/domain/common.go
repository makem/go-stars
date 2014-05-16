package domain

const DateShortFormat = "02.01.2006"

const MaleGenderId = "M"
const FemaleGenderId = "F"

type address struct {
	CountryΩ    ICountry `bson:"country"`
	CityΩ       string   `bson:"city"`
	DistrictΩ   string   `bson:"district"`
	StreetΩ     string   `bson:"street"`
	HouseΩ      string   `bson:"house"`
	BuildingΩ   string   `bson:"building"`
	AppartmentΩ string   `bson:"appartment"`
}

func NewAddress(country ICountry, city, district, street, house, building, appartment string) IAddress {
	return &address{
		country,
		city,
		district,
		street,
		house,
		building,
		appartment,
	}
}

func (address *address) Country(Ω ...ICountry) ICountry {
	if len(Ω) > 0 {
		address.CountryΩ = Ω[0]
	}
	return address.CountryΩ
}

func (address *address) City(city ...string) string {
	if len(city) > 0 {
		address.CityΩ = city[0]
	}
	return address.CityΩ
}

func (address *address) District(district ...string) string {
	if len(district) > 0 {
		address.DistrictΩ = district[0]
	}
	return address.DistrictΩ

}

func (address *address) Street(street ...string) string {
	if len(street) > 0 {
		address.StreetΩ = street[0]
	}
	return address.StreetΩ

}

func (address *address) House(house ...string) string {
	if len(house) > 0 {
		address.HouseΩ = house[0]
	}
	return address.HouseΩ

}

func (address *address) Building(building ...string) string {
	if len(building) > 0 {
		address.BuildingΩ = building[0]
	}
	return address.BuildingΩ

}

func (address *address) Appartment(appartment ...string) string {
	if len(appartment) > 0 {
		address.AppartmentΩ = appartment[0]
	}
	return address.AppartmentΩ

}
