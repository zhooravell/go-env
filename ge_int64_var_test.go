package env

import (
	"os"
	"testing"
)

type int64FlagTest struct {
	valueIn  string
	valueOut int64
	fallback int64
}

var getInt64VarFlagTests = []int64FlagTest{
	{"0", 0, 7},
	{"-0", 0, 7},
	{"-1", -1, 7},
	{"12345", 12345, 7},
	{"-12345", -12345, 7},
	{"-12345", -12345, 7},
	{"012345", 12345, 7},
	{"-012345", -12345, 7},
	{"98765432100", 98765432100, 7},
	{"-98765432100", -98765432100, 7},
	{"9223372036854775807", 1<<63 - 1, 7},
	{"-9223372036854775807", -(1<<63 - 1), 7},
}

var getInt64VarFlagTestsInvalid = []string{
	"test",
	"!!!",
	"+",
	"-",
	"10.5",
}

// Test option with preset value
func TestGetInt64Var(t *testing.T) {
	for _, tt := range getInt64VarFlagTests {
		t.Run(tt.valueIn, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.valueIn); err != nil {
				t.Error(err)
			}

			v, err := GeInt64Var(variableName, tt.fallback)

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
func TestGetInt64VarDefault(t *testing.T) {
	os.Clearenv()

	v, err := GeInt64Var(variableName, 12)

	if err != nil {
		t.Error(err)
	}

	if v != 12 {
		t.Error("Variable must contain 12")
	}
}

// Test option when variable contain invalid value
func TestGetInt64VarInvalidValue(t *testing.T) {
	for _, v := range getInt64VarFlagTestsInvalid {
		t.Run(v, func(t *testing.T) {
			os.Clearenv()
			if err := os.Setenv(variableName, v); err != nil {
				t.Error(err)
			}

			v, err := GeInt64Var(variableName, 0)

			if err == nil {
				t.Error("Must be error")
			}

			if v != 0 {
				t.Error("Variable must contain 0")
			}
		})
	}
}
