package hellonew_test

import (
	"testing"

	"github.com/dmitris/talks/examples/hellonew"
)

func TestGreeting(t *testing.T) {
	want := "Griaß di God!"
	if s := hellonew.LocalGreeting(); s != want {
		t.Errorf("bad local greeting %s, want %s", s, want)
	}
}
