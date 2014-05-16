package regulators_test

import (
	. "domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	r "repository/regulators"
	"storage"
	"testing"
	"time"
)

func TestRegulators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Regulators Repository Suite")
}

const (
	TestDatabaseName string = "stars"
	AssayId                 = "Терапевт"
	AssayName               = "Терапевт"
)

var assays = []IMedicalAssay{
	r.NewMedicalAssay("Кардиолог", "Кардиолог", 12, true),
	r.NewMedicalAssay("Стоматолог", "Стоматолог", 6, false),
	r.NewMedicalAssay("Окулист", "Окулист", 6, true),
	r.NewMedicalAssay("Венеролог", "Венеролог", 12, false),
}

var _ = Describe("Regulators Repository", func() {
	var (
		repository r.IRegulatorsRepository
	)

	BeforeEach(func() {
		repository = r.NewRegulatorsRepository(TestDatabaseName)
	})

	It("should create a new medical assay", func() {
		storage.TruncateCollection(TestDatabaseName, r.AssaysCollectionName)
		assay, err := repository.CreateMedicalAssay(AssayId, AssayName, 6, true)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(assay.Id()).Should(Equal(AssayId))
		Ω(assay.Name()).Should(Equal(AssayName))
		Ω(assay.Period()).Should(Equal(6))
		Ω(assay.Active()).Should(Equal(true))
	})

	It("should read the last saved medical assay", func() {
		list, err := repository.AllMedicalAssays()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(list)).Should(Equal(1))
		Ω(list[0].Id()).Should(Equal(AssayId))
		Ω(list[0].Name()).Should(Equal(AssayName))
		Ω(list[0].Period()).Should(Equal(6))
		Ω(list[0].Active()).Should(Equal(true))
	})

	It("should create another medical assays", func() {
		for _, v := range assays {
			assay, err := repository.CreateMedicalAssay(v.Id(), v.Name(), v.Period(), v.Active())
			Ω(err).ShouldNot(HaveOccurred())
			Ω(assay.Id()).Should(Equal(v.Id()))
			Ω(assay.Name()).Should(Equal(v.Name()))
			Ω(assay.Period()).Should(Equal(v.Period()))
			Ω(assay.Active()).Should(Equal(v.Active()))

		}
	})
	It("should read all the saved medical assays", func() {
		list, err := repository.AllMedicalAssays()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(list)).Should(Equal(5))
	})
	// Assay visits
	const (
		TestLogin = "Petrova"
	)
	It("should get all the visits (phisical & virtual) for the user", func() {
		visits, err := repository.AllEmployeeAssaysLastVisits(TestLogin)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(visits)).Should(Equal(5))
		for _, v := range visits {
			Ω(v.EmployeeLogin()).Should(Equal(TestLogin))
			Ω(v.Visited()).Should(Equal(time.Time{}))
		}
	})

})
