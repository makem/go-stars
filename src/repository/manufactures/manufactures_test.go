package manufactures_test

import (
	//	"fmt"
	. "domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	e "repository/manufactures"
	"storage"
	"testing"
)

func TestManufactures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Manufactures Repository Suite")
}

const (
	TestDatabaseName string = "testdb"
)

var address = NewAddress("Lithuania", "Klaipeda", "", "Rivel", "1", "", "12")
var _ = Describe("Manufacture Repository", func() {
	var (
		repository e.IManufacturesRepository
	)

	BeforeEach(func() {
		repository = e.NewManufacturesRepository(TestDatabaseName)
	})

	Context("Empty database", func() {
		BeforeEach(func() {
			storage.TruncateDatabase(TestDatabaseName)
		})
		It("should create a new manufacture", func() {
			man, err := repository.CreateNew("Klaipeda", address)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(man.Name()).Should(Equal("Klaipeda"))
			Ω(man.Address().Country()).Should(Equal("Lithuania"))
		})
	})
})
