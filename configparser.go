// Copyright (C) 2018 Sebastian Stauch
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package configparser

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/ghodss/yaml"
)

// SetValuesFromEnvironmentTag iterates over the struct and checks for "env" tags. If an "env" tag was found,
// it will set the value of the struct member to the value from the specified env var.
func SetValuesFromEnvironmentTag(config interface{}) {
	t := reflect.Indirect(reflect.ValueOf(config)).Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("env")
		value := os.Getenv(tag)
		if len(value) > 0 {
			// we need more types here
			switch field.Type.Kind().String() {
			case "string":
				reflect.Indirect(reflect.ValueOf(config)).Field(i).SetString(value)
			case "int":
				intval, _ := strconv.Atoi(value)
				reflect.Indirect(reflect.ValueOf(config)).Field(i).SetInt(int64(intval))
			}
		}

		if field.Type.Kind().String() == "struct" {
			s := reflect.Indirect(reflect.ValueOf(config)).Field(i)
			SetValuesFromEnvironmentTag(s.Addr().Interface())
		}
	}
}

// SetValuesFromEnvironment iterates over the struct and checks if environment variables were set.
// It will transform the name of the field to upper and join the the prefix with an _.
// The resulting string will be looked up with os.Getenv if found it will set the value of the struct member
// to the value from the env var.
func SetValuesFromEnvironment(prefix string, config interface{}) {
	if len(prefix) > 0 {
		prefix = prefix + "_"
	}
	t := reflect.Indirect(reflect.ValueOf(config)).Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := strings.ToUpper(field.Name)
		checkvar := prefix + name
		value := os.Getenv(checkvar)
		if len(value) > 0 {
			// we need more types here
			switch field.Type.Kind().String() {
			case "string":
				reflect.Indirect(reflect.ValueOf(config)).Field(i).SetString(value)
			case "int":
				intval, _ := strconv.Atoi(value)
				reflect.Indirect(reflect.ValueOf(config)).Field(i).SetInt(int64(intval))
			}
		}

		if field.Type.Kind().String() == "struct" {
			s := reflect.Indirect(reflect.ValueOf(config)).Field(i)
			SetValuesFromEnvironment(checkvar, s.Addr().Interface())
		}
	}
}

// ParseYaml reads the file at filename and marshals the content to config
func ParseYaml(filename string, config interface{}) error {
	configbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configbytes, config)
	if err != nil {
		return err
	}
	return nil
}

// ParseJSON reads the file at filename and marshals the content to config
func ParseJSON(filename string, config interface{}) error {
	configbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(configbytes, config)
	if err != nil {
		return err
	}
	return nil
}
