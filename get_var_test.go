package env

import (
	"os"
	"testing"
)

type flagTest struct {
	value    string
	fallback string
}

var (
	variableName    = "OS_ENV_TEST_VARIABLE"
	getVarFlagTests = []flagTest{
		{"google.com", "www.bing.com"},
		{"cat", "dog"},
		{"test", "dev"},
	}
)

// Test option with preset value
func TestGetVar(t *testing.T) {
	for _, tt := range getVarFlagTests {
		t.Run(tt.value, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.value); err != nil {
				t.Error(err)
			}

			if GetVar(variableName, tt.fallback) != tt.value {
				t.Errorf("Variable %s not equal to value %s", variableName, tt.value)
			}
		})
	}
}

// Test option when variable was not found
func TestGetVarDefault(t *testing.T) {
	os.Clearenv()
	defaultValue := "--TEST--"

	if GetVar(variableName, defaultValue) != defaultValue {
		t.Errorf("Variable %s not equal to default value %s", variableName, defaultValue)
	}
}
