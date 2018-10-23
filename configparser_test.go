package configparser

import (
	"os"
	"testing"
)

type Address struct {
	Street string `json:"street" env:"USER_ADDRESS_STREET"`
	City   string `json:"city" env:"USER_ADDRESS_CITY"`
}

type User struct {
	Name    string  `json:"name" env:"USER_NAME"`
	Age     int     `json:"age" env:"USER_AGE"`
	Species string  `json:"species" env:"USER_SPECIES"`
	Street  string  `json:"street" env:"USER_STREET"`
	NoEnv   string  `json:"noenv"`
	Address Address `json:"address"`
}

func TestParseYaml(t *testing.T) {
	var user User
	err := ParseYaml("test.yaml", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if user.Species != "unknown" {

		t.Errorf("user.species is not unknown.")

	}
}

func TestSetEnv(t *testing.T) {
	var user User
	err := ParseYaml("test.yaml", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	os.Setenv("USER_NAME", "thanos")
	os.Setenv("USER_ADDRESS_STREET", "newstreet")
	SetValuesFromEnvironment(&user)
	if user.Name != "thanos" {
		t.Errorf("Name was not set from env var.")
	}

	if user.Address.Street != "newstreet" {
		t.Errorf("Street was not set in Address.")
	}

	t.Logf("%#v", user)

}
