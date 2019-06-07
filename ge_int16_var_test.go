package env

import (
	"os"
	"testing"
)

type int16FlagTest struct {
	valueIn  string
	valueOut int16
	fallback int16
}

var getInt16VarFlagTests = []int16FlagTest{
	{"0", 0, 7},
	{"-0", 0, 7},
	{"1", 1, 7},
	{"-1", -1, 7},
	{"-50", -50, 7},
	{"50", 50, 7},
	{"-128", -128, 7},
	{"127", 127, 7},
	{"-32768", -32768, 7},
	{"32767", 32767, 7},
}

var getInt16VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
	"2147483649",
}

// Test option with preset value
func TestGetInt16Var(t *testing.T) {
	for _, tt := range getInt16VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GeInt16Var(variableName, tt.fallback)

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
func TestGetInt16VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GeInt16Var(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetInt16VarInvalidValue(t *testing.T) {
	for _, v := range getInt16VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GeInt16Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
