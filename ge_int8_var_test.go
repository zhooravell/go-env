package env

import (
	"os"
	"testing"
)

type int8FlagTest struct {
	valueIn  string
	valueOut int8
	fallback int8
}

var getInt8VarFlagTests = []int8FlagTest{
	{"0", 0, 7},
	{"-0", 0, 7},
	{"1", 1, 7},
	{"-1", -1, 7},
	{"-50", -50, 7},
	{"50", 50, 7},
	{"-128", -128, 7},
	{"127", 127, 7},
}

var getInt8VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
	"2147483649",
}

// Test option with preset value
func TestGetInt8Var(t *testing.T) {
	for _, tt := range getInt8VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GeInt8Var(variableName, tt.fallback)

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
func TestGetInt8VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GeInt8Var(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetInt8VarInvalidValue(t *testing.T) {
	for _, v := range getInt8VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GeInt8Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
