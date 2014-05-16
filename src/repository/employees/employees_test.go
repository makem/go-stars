package employees_test

import (
	. "domain"
	"fmt"
	//. "github.com/ahmetalpbalkan/go-linq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	e "repository/employees"
	g "repository/geography"
	"storage"
	"testing"
	"time"
)

func TestEmployee(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Employee Repository Suite")
}

const (
	TestDatabaseName string = "stars"
)

var _ = Describe("Employee Repository", func() {
	var (
		repository e.IEmployeeRepository
		geo        g.IGeographyRepository
	)

	BeforeEach(func() {
		repository = e.NewEmployeeRepository(TestDatabaseName)
		geo = g.NewGeographyRepository(TestDatabaseName)
		//storage.TruncateDatabase(TestDatabaseName)
	})

	XIt("should return employer by login", func() {
		var employeeLogin = "Ivanov"
		e, err := repository.FindByLogin(employeeLogin)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(e.Login()).Should(Equal(employeeLogin))
		Expect(e.Name()).Should(Equal("Name"))
		fmt.Println(e)
	})

	XIt("should return all the employee roles", func() {
		roles := repository.AllEmployeeRoles()
		Ω(len(roles)).Should(Equal(5))
	})

	It("should", func() {
		var roles []string
		if len(roles) == 0 {
			roles = append(roles, "hello")
		}
		Ω(len(roles)).Should(Equal(1))
	})

	It("should create a new employee", func() {
		storage.TruncateCollection(TestDatabaseName, "employees")
		var employeeLogin = "Petrova"
		emp, err := repository.CreateNew("Петрова", "Светлана", "Валерьевна", employeeLogin, "12345", nil)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(emp.Name()).Should(Equal("Петрова Светлана"))
		Ω(emp.Login()).Should(Equal(employeeLogin))
		Ω(emp.Password()).Should(Equal("12345"))
		emp.AssignRoleIds(repository.AllEmployeeRoles()[0].Id())
		emp.AssignRoleIds(repository.AllEmployeeRoles()[1].Id())
		emp.GenderId(FemaleGenderId)
		emp.CountryId("LT")
		repository.SaveEmployee(emp)
		emp2, err2 := repository.FindByLogin(employeeLogin)
		Ω(err2).ShouldNot(HaveOccurred())
		Ω(len(emp2.RoleIds())).Should(Equal(2))
		Ω(emp2.RoleIds()[0]).Should(Equal("ADM"))
		Ω(emp2.RoleIds()[1]).Should(Equal("DSP"))

		employeeLogin = "Pavel"
		emp, err = repository.CreateNew("Выговский", "Павел", "Михайлович", employeeLogin, "12345", nil)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(emp.Name()).Should(Equal("Выговский Павел"))
		Ω(emp.Login()).Should(Equal(employeeLogin))
		Ω(emp.Password()).Should(Equal("12345"))
		emp.AssignRoleIds(repository.AllEmployeeRoles()[0].Id())
		emp.Active(true)
		emp.GenderId(MaleGenderId)
		emp.CountryId("LT")
		repository.SaveEmployee(emp)
		emp2, err2 = repository.FindByLogin(employeeLogin)
		Ω(err2).ShouldNot(HaveOccurred())
		Ω(len(emp2.RoleIds())).Should(Equal(1))
		Ω(emp2.RoleIds()[0]).Should(Equal("ADM"))
		Ω(emp2.Active()).Should(BeTrue())

	})

	XIt("should update employer", func() {
		var employeeLogin = "Employee2"
		e, err := repository.FindByLogin(employeeLogin)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(e.Login()).Should(Equal(employeeLogin))
		e.Active(true)
		err = repository.SaveEmployee(e)
		Expect(err).ShouldNot(HaveOccurred())
		//fmt.Println(e)
	})

	XIt("should list existing employeers", func() {
		start := time.Now()
		list, err := repository.AllEmployees()
		elapsed := time.Since(start)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(len(list)).Should(Equal(1))
		fmt.Printf("Repository call took: %s", elapsed)
	})

	It("should authenticate employeers", func() {
		start := time.Now()
		emp, err := repository.AuthenticateEmployee("Petrova", "12345")
		elapsed := time.Since(start)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(emp.Login()).Should(Equal("Employee2"))
		fmt.Printf("Repository call took: %s", elapsed)
	})

	It("should provide role assign operations", func() {
		emp := e.NewEmployee("", "", "", "Test", nil)
		emp.AssignRoleIds("1", "2", "3")
		Ω(len(emp.RoleIds())).Should(Equal(3))
		emp.AssignRoleIds("4", "1", "5")
		Ω(len(emp.RoleIds())).Should(Equal(5))
		emp.AssignRoleIds("1", "2", "3", "4", "5")
		Ω(len(emp.RoleIds())).Should(Equal(5))
	})
	It("should provide role reassign operations", func() {
		emp := e.NewEmployee("", "", "", "Test", nil)
		emp.AssignRoleIds("1", "2", "3")
		Ω(len(emp.RoleIds())).Should(Equal(3))
		emp.ReAssignRoleIds("4", "1", "5")
		Ω(len(emp.RoleIds())).Should(Equal(3))
		emp.AssignRoleIds("2", "3")
		Ω(len(emp.RoleIds())).Should(Equal(5))

	})

	It("should complete name", func() {
		emp := e.NewEmployee("Last", "First", "Middle", "Test", nil)
		Ω(emp.Name()).Should(Equal("Last First"))
	})

})
