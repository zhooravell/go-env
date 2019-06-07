package env

import (
	"os"
	"testing"
)

type uintFlagTest struct {
	valueIn  string
	valueOut uint
	fallback uint
}

var getUintVarFlagTests = []uintFlagTest{
	{"0", 0, 7},
	{"42", 42, 7},
	{"13", 13, 7},
	{"432", 432, 7},
	{"1024", 1024, 7},
	{"1989", 1989, 7},
	{"20001", 20001, 7},
}

var getUintVarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
}

// Test option with preset value
func TestGetUintVar(t *testing.T) {
	for _, tt := range getUintVarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GeUintVar(variableName, tt.fallback)

			if err != nil {
				t.Error(err)
			}

			if v != tt.valueOut {
				t.Errorf("Variable %s not equal to value '%v'", variableName, tt.valueOut)
			}
		})
	}
}

// Test option when variable was not found
func TestGetUintVarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GeUintVar(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetUintVarInvalidValue(t *testing.T) {
	for _, v := range getUintVarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GeUintVar(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
