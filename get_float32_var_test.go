package env

import (
	"os"
	"testing"
)

type float32FlagTest struct {
	valueIn  string
	valueOut float32
	fallback float32
}

var getFloat32VarFlagTests = []float32FlagTest{
	{"1.000000059604644775390625", 1, 7},
	{"1.000000059604644775390626", 1.0000001, 7},
	{"4951760157141521099596496896", 4.9517602e+27, 7},
	{"3.402823567e38", 3.4028235e+38, 7},
	{"-3.402823567e38", -3.4028235e+38, 7},
}

var getFloat32VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"3.4028235678e38",
	"1.797693134862315808e308",
}

// Test option with preset value
func TestGetFloat32Var(t *testing.T) {
	for _, tt := range getFloat32VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GetFloat32Var(variableName, tt.fallback)

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
func TestGetFloat32VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GetFloat32Var(variableName, 1.001)

	if err != nil {
		t.Error(err)
	}

	if v != 1.001 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetFloat32VarInvalidValue(t *testing.T) {
	for _, v := range getFloat32VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GetFloat32Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
