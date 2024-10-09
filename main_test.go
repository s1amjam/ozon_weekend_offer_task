package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	td := &testData{}

	td.add()

	if len(td.phones) != 1 {
		t.Errorf("Expected length 1, got %d", len(td.phones))
	}
}

func TestGenerate(t *testing.T) {
	td := &testData{}
	n := 100
	generate(n, td)

	if len(td.phones) != n {
		t.Errorf("Expected %d phone numbers, but got %d", n, len(td.phones))
	}
}

func TestPhoneFormat(t *testing.T) {
	td := &testData{}

	for i := 0; i < 1000; i++ {
		phone := td.randPhone()

		if phone < 89000000000 || phone > 89999999999 {
			t.Errorf("Wrong phone format: %d", phone)
		}
	}
}

func TestGenerateWithZero(t *testing.T) {
	td := &testData{}
	n := 0
	generate(n, td)

	if len(td.phones) != 0 {
		t.Errorf("Expected 0 but got %d", len(td.phones))
	}
}

func TestGenerateWithNegativeNumber(t *testing.T) {
	td := &testData{}
	n := -100
	generate(n, td)

	if len(td.phones) != 0 {
		t.Errorf("Expected 0 while trying to generate with negative number, but got %d", len(td.phones))
	}
}
