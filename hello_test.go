package hello

import (
	"testing"
	"github.com/stretchr/testify/assert"
	."github.com/onsi/ginkgo"
	."github.com/onsi/gomega"
)

// Librería estandar de Golang
func TestHello(t *testing.T) {
	want := "Hola Mundo."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}


func TestSomething(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  
  }
func TestCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("Calculator", func() {
	Describe("Add numbers", func() {
		Context("1 and 2", func() {
			It("should be 3", func() {
				Expect(Add(1, 2)).To(Equal(3))
			})
		})
	})

	Describe("Subtract numbers", func() {
		Context("3 from 5", func() {
			It("should be 2", func() {
				Expect(Subtract(5, 3)).To(Equal(2))
			})
		})
	})

	Describe("Multiply numbers", func() {
		Context("5 with 6", func() {
			It("should be 30", func() {
				Expect(Multiply(5, 6)).To(Equal(30))
			})
		})
	})

	Describe("Divide numbers", func() {
		Context("10 by 2", func() {
			It("should be 30", func() {
				Expect(Divide(10, 2)).To(Equal(float64(5)))
			})
		})
	})
})