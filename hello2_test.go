package hello

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}


func TestSomething(t *testing.T) {

	var a string = "Heo"
	var b string = "Hello"
  
	assert.Equal(t, a, b, "The two words should be the same.")
  
}