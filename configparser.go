// Copyright (C) 2018 Sebastian Stauch
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package configparser

import (
	"encoding/json"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

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
