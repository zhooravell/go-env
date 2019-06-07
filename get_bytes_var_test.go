package env

import (
	"bytes"
	"os"
	"testing"
)

type bytesFlagTest struct {
	value    string
	fallback []byte
}

var getVarBytesFlagTests = []bytesFlagTest{
	{"google.com", []byte("www.bing.com")},
	{"cat", []byte("dog")},
	{"test", []byte("dev")},
}

// Test option with preset value
func TestGetBytesVar(t *testing.T) {
	for _, tt := range getVarBytesFlagTests {
		t.Run(tt.value, func(t *testing.T) {
			os.Clearenv()

			if err := os.Setenv(variableName, tt.value); err != nil {
				t.Error(err)
			}

			if !bytes.Equal([]byte(tt.value), GetBytesVar(variableName, tt.fallback)) {
				t.Errorf("Variable %s not equal to value %s", variableName, tt.value)
			}
		})
	}
}

// Test option when variable was not found
func TestGetBytesVarDefault(t *testing.T) {
	os.Clearenv()
	defaultValue := []byte("--TEST--")

	if !bytes.Equal(defaultValue, GetBytesVar(variableName, defaultValue)) {
		t.Errorf("Variable %s not equal to default value %s", variableName, defaultValue)
	}
}
