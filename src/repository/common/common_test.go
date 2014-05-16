package common_test

import (
	. "domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	c "repository/common"
	er "repository/employees"
	"testing"
	"time"
)

func TestCommonDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Common Domain Suite")
}

const (
	NAME     = "name"
	ID       = "12345"
	ACTIVE   = true
	INACTIVE = false
)

var _ = Describe("Common domain implementation", func() {
	It("can be unique", func() {
		u := c.Unique{IdΩ: ID}
		iu := IUnique(u)
		var iu2 IUnique = u
		Ω(iu.Id()).Should(Equal(ID))
		Ω(iu2.Id()).Should(Equal(ID))
	})
	It("can be named", func() {
		v := c.Named{NameΩ: NAME}
		vI := INamed(&v)
		var vII INamed = &v
		Ω(vI.Name()).Should(Equal(NAME))
		Ω(vII.Name()).Should(Equal(NAME))
		//Assigment
		vI.Name(NAME + "/")
		Ω(vII.Name()).Should(Equal(NAME + "/"))
	})
	It("can be active", func() {
		v := c.Activable{ActiveΩ: ACTIVE}
		vI := IActivable(&v)
		var vII IActivable = &v
		Ω(vI.Active()).Should(Equal(ACTIVE))
		Ω(vII.Active()).Should(Equal(ACTIVE))
		//Assigment
		vI.Active(INACTIVE)
		Ω(vII.Active()).Should(Equal(INACTIVE))

	})
	It("can be tracked", func() {
		employee := er.NewEmployee("", "", nil)
		now := time.Now()
		v := c.Trackable{
			CreatedΩ:   now,
			CreatedByΩ: employee,
		}
		vI := ITrackable(&v)
		Ω(vI.Created()).Should(Equal(now))
		Ω(vI.CreatedBy()).Should(Equal(employee))

	})
})
