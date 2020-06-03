package tests

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted: %q, got: %q", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Goodbye, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected result. Wanted: %q, got: %q", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct {
		name   string
		expect string
	}{
		{name: "Gopher", expect: "Hello, Gopher!\n"},
		{name: "", expect: "Hello, !\n"},
	}
	for _, s := range scenarios {
		got := Greet(s.name)
		if got != s.expect {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got %q", s.name, got, s.expect)
		}
	}
}
