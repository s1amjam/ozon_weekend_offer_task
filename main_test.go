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
	var tests = []struct {
		name   string
		input  int
		expect int
	}{
		{"Input 100", 100, 100},
		{"Input 0", 0, 0},
		{"Input -100", -100, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			td := &testData{}
			generate(tt.input, td)

			if len(td.phones) != tt.expect {
				t.Errorf("Expected %d phone numbers, but got %d", tt.expect, len(td.phones))
			}
		})
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
