package env

import (
	"os"
	"testing"
)

type int32FlagTest struct {
	valueIn  string
	valueOut int32
	fallback int32
}

var getInt32VarFlagTests = []int32FlagTest{
	{"0", 0, 7},
	{"-0", 0, 7},
	{"1", 1, 7},
	{"-1", -1, 7},
	{"12345", 12345, 7},
	{"-12345", -12345, 7},
	{"012345", 12345, 7},
	{"-012345", -12345, 7},
	{"987654321", 987654321, 7},
	{"-987654321", -987654321, 7},
	{"2147483647", 1<<31 - 1, 7},
	{"-2147483647", -(1<<31 - 1), 7},
}

var getInt32VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
	"2147483649",
}

// Test option with preset value
func TestGetInt32Var(t *testing.T) {
	for _, tt := range getInt32VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GeInt32Var(variableName, tt.fallback)

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
func TestGetInt32VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GeInt32Var(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetInt32VarInvalidValue(t *testing.T) {
	for _, v := range getInt32VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GeInt32Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
