package env

import (
	"os"
	"testing"
)

type uint8FlagTest struct {
	valueIn  string
	valueOut uint8
	fallback uint8
}

var getUint8VarFlagTests = []uint8FlagTest{
	{"0", 0, 7},
	{"42", 42, 7},
	{"13", 13, 7},
	{"255", 255, 7},
}

var getUint8VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
	"-1",
}

// Test option with preset value
func TestGetUint8Var(t *testing.T) {
	for _, tt := range getUint8VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GetUint8Var(variableName, tt.fallback)

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
func TestGetUint8VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GetUint8Var(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetUint8VarInvalidValue(t *testing.T) {
	for _, v := range getUint8VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GetUint8Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
