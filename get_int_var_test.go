package env

import (
	"os"
	"testing"
)

type intFlagTest struct {
	valueIn  string
	valueOut int
	fallback int
}

var getIntVarFlagTests = []intFlagTest{
	{"-42", -42, 11},
	{"13", 13, 11},
	{"432", 432, 11},
	{"1024", 1024, 11},
	{"1989", 1989, 11},
	{"20001", 20001, 11},
}

var getIntVarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
}

// Test option with preset value
func TestGetIntVar(t *testing.T) {
	for _, tt := range getIntVarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GetIntVar(variableName, tt.fallback)

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
func TestGetIntVarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GetIntVar(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetIntVarInvalidValue(t *testing.T) {
	for _, v := range getIntVarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GetIntVar(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
