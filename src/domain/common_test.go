package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSlice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Employee Repository Suite")
}

var _ = Describe("Slices investigation", func() {

	BeforeEach(func() {

	})

	It("Резервирование пространства", func() {
		s := make([]int, 0, 10)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(10))
	})

	It("Расширение слайса без превышения порога", func() {
		s := make([]int, 0, 10)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(10))
		s = append(s, 5, 3, 8)
		Ω(len(s)).Should(Equal(3))
		Ω(cap(s)).Should(Equal(10))
	})

	It("Расширение слайса до порога включительно (3)", func() {
		s := make([]int, 0, 3)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(3))
		s = append(s, 5, 3, 8)
		Ω(len(s)).Should(Equal(3))
		Ω(cap(s)).Should(Equal(3))
	})

	It("Перенасыщение слайса с выделением дополнительного сегмента", func() {
		s := make([]int, 0, 3)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(3))
		s = append(s, 5, 3, 8, 4, 6)
		Ω(len(s)).Should(Equal(5))
		Ω(cap(s)).Should(Equal(8))
	})

	It("Расширение слайса до порога включительно (5)", func() {
		s := make([]int, 0, 5)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(5))
		s = append(s, 5, 3, 8, 4, 6)
		Ω(len(s)).Should(Equal(5))
		Ω(cap(s)).Should(Equal(5))
	})
	It("Расширение слайса с 3 при вставке 10-ти приводит с резервации 10-ти ", func() {
		s := make([]int, 0, 3)
		Ω(len(s)).Should(Equal(0))
		Ω(cap(s)).Should(Equal(3))
		v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		s = append(s, v...)
		Ω(len(s)).Should(Equal(10))
		Ω(cap(s)).Should(Equal(10))
	})

	// changing
	It("Реслайсинг приводит к дублированю и копированию элементов массива", func() {
		s := make([]int, 0, 3)
		s = append(s, 1, 2, 3)
		v := append(s, 4, 5, 6)
		Ω(len(s)).Should(Equal(3))
		Ω(len(v)).Should(Equal(6))
		s[0] = 7
		Ω(s[0]).Should(Equal(7))
		Ω(v[0]).Should(Equal(1)) // Массив был скопирован

	})

	// changing
	It("Расширение без реслайсинга не приводит к копированию массива", func() {
		s := make([]int, 0, 3)
		s = append(s, 1, 2)
		v := append(s, 3)
		Ω(len(s)).Should(Equal(2))
		Ω(len(v)).Should(Equal(3))
		s[0] = 7
		Ω(s[0]).Should(Equal(7))
		Ω(v[0]).Should(Equal(7)) //Массив не был скопирован

	})
})
