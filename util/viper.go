package util

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BindEnvsIntoViper(v *viper.Viper, obj any, parentPrefix string) {
	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			panic("obj is a nil pointer")
		}
		val = val.Elem()
	}

	if !val.IsValid() {
		panic("invalid reflect.Value")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		rawTag := field.Tag.Get("mapstructure")
		if rawTag == "" {
			continue
		}

		parts := strings.Split(rawTag, ",")
		tag := parts[0]
		hasOmitempty := false
		for _, p := range parts[1:] {
			if strings.TrimSpace(p) == "omitempty" {
				hasOmitempty = true
				break
			}
		}

		fieldType := field.Type
		fieldVal := val.Field(i)

		// Recurse into nested structs
		nestedPrefix := tag
		if parentPrefix != "" {
			nestedPrefix = parentPrefix + "_" + tag
		}

		// Handle pointer-to-struct
		if fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct {
			var target reflect.Value
			if fieldVal.IsNil() {
				target = reflect.New(fieldType.Elem()).Elem() // create dummy value
			} else {
				target = fieldVal.Elem()
			}
			if !hasOmitempty {
				BindEnvsIntoViper(v, target.Addr().Interface(), nestedPrefix)
			}
			continue
		} else if fieldType.Kind() == reflect.Struct {
			// Handle direct struct
			if !hasOmitempty {
				BindEnvsIntoViper(v, fieldVal.Addr().Interface(), nestedPrefix)
			}
			continue
		}

		// Construct env var name (flat)
		envKey := tag
		if parentPrefix != "" {
			envKey = parentPrefix + "_" + tag
		}
		envKey = strings.ToUpper(envKey)

		// Construct nested Viper key
		nestedKey := tag
		if parentPrefix != "" {
			nestedKey = parentPrefix + "." + tag
		}

		// Enforce required if no omitempty
		if !hasOmitempty && os.Getenv(envKey) == "" {
			panic(fmt.Errorf("missing required environment variable: %s", envKey))
		}

		// Bind env var
		err := v.BindEnv(nestedKey, envKey)
		cobra.CheckErr(err)

		// Inject into nested Viper key
		v.Set(nestedKey, v.Get(envKey))
	}
}