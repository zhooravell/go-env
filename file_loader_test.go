package env

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Clearenv()
	if err := os.Setenv("ENV_E", "overload"); err != nil {
		t.Error(err)
	}

	envContent := []byte(`
ENV_A=1
ENV_B='https://google.com'
ENV_C=c
ENV_D="test\n test"
ENV_E=2
ENV_F="test\r test"
ENV_G="test=test\\a"
`)

	if err := ioutil.WriteFile(defaultFileName, envContent, 0777); err != nil {
		t.Error(err)
	}

	if err := Load(); err != nil {
		t.Error(err)
	}

	if a := os.Getenv("ENV_A"); a != "1" {
		t.Error("ENV_A must equal 1")
	}

	if b := os.Getenv("ENV_B"); b != "https://google.com" {
		t.Error("ENV_B must equal https://google.com")
	}

	if c := os.Getenv("ENV_C"); c != "c" {
		t.Error("ENV_C must equal 'c'")
	}

	if d := os.Getenv("ENV_D"); d != "test\n test" {
		t.Error("ENV_D must equal 'test\n test'")
	}

	if d := os.Getenv("ENV_F"); d != "test\r test" {
		t.Error("ENV_F must equal 'test\r test'")
	}

	if e := os.Getenv("ENV_E"); e != "overload" {
		t.Error("ENV_E must equal 'overload'")
	}

	if e := os.Getenv("ENV_G"); e != "test=test\\a" {
		t.Error("ENV_G must equal 'test=test\\a")
	}

	if err := os.Remove(defaultFileName); err != nil {
		t.Error(err)
	}
}

func TestLoadWithOverriding(t *testing.T) {
	os.Clearenv()
	if err := os.Setenv("ENV_E", "overload"); err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile("first.env", []byte("ENV_A=1\nENV_B='https://google.com'\n"), 0777); err != nil {
		t.Error(err)
	}

	if err := ioutil.WriteFile("second.env", []byte("ENV_C=c\nENV_E=2\nexport ENV_F=export\n"), 0777); err != nil {
		t.Error(err)
	}

	if err := LoadWithOverriding("first.env", "second.env"); err != nil {
		t.Error(err)
	}

	if a := os.Getenv("ENV_A"); a != "1" {
		t.Error("ENV_A must equal 1")
	}

	if b := os.Getenv("ENV_B"); b != "https://google.com" {
		t.Error("ENV_B must equal https://google.com")
	}

	if c := os.Getenv("ENV_C"); c != "c" {
		t.Error("ENV_C must equal 'c'")
	}

	if e := os.Getenv("ENV_E"); e != "2" {
		t.Error("ENV_E must equal '2'")
	}

	if e := os.Getenv("ENV_F"); e != "export" {
		t.Error("ENV_F must equal 'export'")
	}

	if err := os.Remove("first.env"); err != nil {
		t.Error(err)
	}

	if err := os.Remove("second.env"); err != nil {
		t.Error(err)
	}
}

func TestLoadFileNotFound(t *testing.T) {
	os.Clearenv()
	if err := Load("test.env"); err == nil {
		t.Error("Must been error")
	}
}

func TestLoadWithOverridingFileNotFound(t *testing.T) {
	os.Clearenv()
	if err := LoadWithOverriding("test.env"); err == nil {
		t.Error("Must been error")
	}
}

func TestLoadInvalidFormat(t *testing.T) {
	os.Clearenv()
	if err := ioutil.WriteFile(defaultFileName, []byte("ENV_A"), 0777); err != nil {
		t.Error(err)
	}

	if err := Load(); err == nil {
		t.Error("Must been error")
	}

	if err := os.Remove(defaultFileName); err != nil {
		t.Error(err)
	}
}
