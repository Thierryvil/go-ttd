package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Sum(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("result: '%d', expect: '%d'", sum, expected)
	}
}
