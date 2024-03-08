package util_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/muulinCorp/interlib/util"
)

func TestGetFromEnv(t *testing.T) {
	type testStruct struct {
		StringField string        `env:"STRING_ENV_VAR"`
		IntField    int           `env:"INT_ENV_VAR"`
		BoolField   bool          `env:"BOOL_ENV_VAR"`
		DurField    time.Duration `env:"DURATION_ENV_VAR"`
	}

	// Test for obj not being a pointer
	err := util.GetFromEnv(testStruct{})
	if err == nil || err.Error() != "obj must be a pointer" {
		t.Error("Expected error for non-pointer obj")
	}

	// Test for missing environmental variable
	err = util.GetFromEnv(&testStruct{})
	if err == nil || !strings.Contains(err.Error(), "must not be blank") {
		t.Error("Expected error for missing environmental variable")
	}
	os.Setenv("STRING_ENV_VAR", "string")
	// Test for non-integer environmental variable
	os.Setenv("INT_ENV_VAR", "not_an_integer")
	err = util.GetFromEnv(&testStruct{})
	if err == nil || !strings.Contains(err.Error(), "must be an integer") {
		t.Error("Expected error for non-integer environmental variable")
	}
	// os.Unsetenv("INT_ENV_VAR")
	os.Setenv("INT_ENV_VAR", "1")

	// Test for non-boolean environmental variable
	os.Setenv("BOOL_ENV_VAR", "not_a_boolean")
	err = util.GetFromEnv(&testStruct{})
	if err == nil || !strings.Contains(err.Error(), "must be a boolean") {
		t.Error("Expected error for non-boolean environmental variable")
	}
	os.Setenv("BOOL_ENV_VAR", "true")

	// Test for non-duration environmental variable
	os.Setenv("DURATION_ENV_VAR", "not_a_duration")
	err = util.GetFromEnv(&testStruct{})
	if err == nil || !strings.Contains(err.Error(), "must be a duration") {
		t.Error("Expected error for non-duration environmental variable")
	}
	os.Unsetenv("STRING_ENV_VAR")
	os.Unsetenv("INT_ENV_VAR")
	os.Unsetenv("BOOL_ENV_VAR")
	os.Unsetenv("DURATION_ENV_VAR")

	// Test for unsupported type environmental variable
	os.Setenv("MAP_ENV_VAR", "1")
	type unsupportedTypeStruct struct {
		MapField map[string]string `env:"MAP_ENV_VAR"`
	}
	err = util.GetFromEnv(&unsupportedTypeStruct{})
	if err == nil || !strings.Contains(err.Error(), "unsupported type") {
		t.Error("Expected error for unsupported type environmental variable")
	}
}
