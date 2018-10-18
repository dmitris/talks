package hellomun_test

import (
	"testing"

	"github.com/dmitris/talks/examples/hellomun"
)

func TestGreeting(t *testing.T) {
	tests := []struct {
		in, want, label string
	}{
		{"gopher", "Servus, gopher!", "non-empty input"},
		{"", "Servus, world!", "empty input"},
	}
	for _, tt := range tests {
		if s := hellomun.Greeting(tt.in); s != tt.want {
			t.Errorf("%s - bad greeting %s, want %s", tt.label, s, tt.want)
		}
	}
}
