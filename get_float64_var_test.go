package env

import (
	"os"
	"testing"
)

type float64FlagTest struct {
	valueIn  string
	valueOut float64
	fallback float64
}

var getFloat64VarFlagTests = []float64FlagTest{
	{"99999999999999974834176", 9.999999999999997e+22, 7},
	{"-0.1", -0.1, 7},
	{"1.7976931348623157e308", 1.7976931348623157e+308, 7},
	{"-1.7976931348623158e308", -1.7976931348623157e+308, 7},
	{"22.222222222222222", 22.222222222222222, 7},
}

var getFloat64VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
}

// Test option with preset value
func TestGetFloat64Var(t *testing.T) {
	for _, tt := range getFloat64VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GetFloat64Var(variableName, tt.fallback)

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
func TestGetFloat64VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GetFloat64Var(variableName, 1.001)

	if err != nil {
		t.Error(err)
	}

	if v != 1.001 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetFloat64VarInvalidValue(t *testing.T) {
	for _, v := range getFloat64VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GetFloat64Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
