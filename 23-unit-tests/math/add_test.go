package math

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}

	_, err2 := Add(1, -2)
	if err2 == nil {
		t.Error("second arg negative - expected error not be nil")
	}

	_, err3 := Add(-1, -2)
	if err3 == nil {
		t.Error("both args negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 1)
	if err == nil {
		t.Error("first arg zero - expected error not be nil")
	}

	_, err2 := Add(1, 0)
	if err2 == nil {
		t.Error("second arg zero - expected error not be nil")
	}
}
