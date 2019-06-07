package env

import (
	"os"
	"testing"
)

type boolFlagTest struct {
	valueIn  string
	valueOut bool
	fallback bool
}

var getBoolVarFlagTests = []boolFlagTest{
	{"true", true, false},
	{"1", true, false},
	{"t", true, false},
	{"T", true, false},
	{"TRUE", true, false},
	{"True", true, false},

	{"0", false, true},
	{"f", false, true},
	{"F", false, true},
	{"false", false, true},
	{"FALSE", false, true},
	{"False", false, true},
}

var getBoolVarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+false",
	"0true",
	"true1",
	"+",
	"-",
	"o",
}

// Test option with preset value
func TestGetBoolVar(t *testing.T) {
	for _, tt := range getBoolVarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GetBoolVar(variableName, tt.fallback)

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
func TestGetBoolVarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GetBoolVar(variableName, false)

	if err != nil {
		t.Error(err)
	}

	if v == true {
		t.Error("Variable must contain false value")
	}
}

// Test option when variable contain invalid value
func TestGetBoolVarInvalidValue(t *testing.T) {
	for _, v := range getBoolVarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GetBoolVar(variableName, false)

			if err == nil {
				t.Error("Must be error")
			}

			if v != false {
				t.Error("Variable must contain false value")
			}
		})
	}
}
