package geography_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	e "repository/geography"
	"storage"
	"testing"
)

func TestGeography(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geography Repository Suite")
}

const (
	TestDatabaseName string = "stars"
	UACode           string = "UA"
	UAName           string = "Ukraine"
	LTCode           string = "LT"
	LTName           string = "Lithuvania"
	RUCode           string = "RU"
	RUName           string = "Russia"
)

var _ = Describe("Geography Repository", func() {
	var (
		repository e.IGeographyRepository
	)

	BeforeEach(func() {
		repository = e.NewGeographyRepository(TestDatabaseName)
	})

	It("should create a new country on empty collection", func() {
		storage.TruncateCollection(TestDatabaseName, e.CountryCollectionName)
		country, err := repository.NewCountry(UACode, UAName)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(country.Id()).Should(Equal(UACode))
		Ω(country.Name()).Should(Equal(UAName))
		list, err := repository.AllCountries()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(list)).Should(Equal(1))
		Ω(list[0].Id()).Should(Equal(UACode))
		Ω(list[0].Name()).Should(Equal(UAName))
	})

	It("should avoid creating a duplicated country", func() {
		_, err := repository.NewCountry(UACode, UAName)
		Ω(err).Should(HaveOccurred())
		Ω(err.Error()).Should(Equal("The country with code 'UA' already exists"))
	})

	It("should create another country", func() {
		country, err := repository.NewCountry(LTCode, LTName)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(country.Id()).Should(Equal(LTCode))
		Ω(country.Name()).Should(Equal(LTName))
		country, err = repository.NewCountry(RUCode, RUName)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(country.Id()).Should(Equal(RUCode))
		Ω(country.Name()).Should(Equal(RUName))
	})

	It("should get all the created countries", func() {
		list, err := repository.AllCountries()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(list)).Should(Equal(3))
		Ω(list[0].Id()).Should(Equal(UACode))
		Ω(list[1].Id()).Should(Equal(LTCode))
		Ω(list[2].Id()).Should(Equal(RUCode))

	})
})
