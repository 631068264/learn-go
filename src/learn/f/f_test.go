package f

import (
	"fmt"
	"testing"
)

var a string

func n() { fmt.Println(a) }

func m() {
	a := "O"
	fmt.Println(a)
}
func TestA(t *testing.T) {
	n()
	m()
	n()
}
func TestB(t *testing.T) {
	n()
	m()
	n()
}
