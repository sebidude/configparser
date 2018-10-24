package configparser

import (
	"os"
	"testing"
)

type Address struct {
	Street string `json:"street" env:"USER_ADDRESS_STREET"`
	City   string `json:"city"`
}

type User struct {
	Name    string  `json:"name" env:"USER_NAME"`
	Age     int     `json:"age" env:"USER_AGE"`
	Species string  `json:"species" env:"USER_SPECIES"`
	Street  string  `json:"street" env:"USER_STREET"`
	NoEnv   string  `json:"noenv"`
	Address Address `json:"address"`
	Limit   int     `json:"limit"`
	CanTalk bool    `json:"cantalk" env:"USER_TALK"`
	CanWalk bool    `json:"canwalk"`
	Score   float64 `json:"score"`
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

func TestParseJSON(t *testing.T) {
	var user User
	err := ParseJSON("test.json", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if user.Species != "unknown" {
		t.Errorf("user.species is not unknown.")
	}
}

func TestSetEnvTag(t *testing.T) {
	var user User
	err := ParseYaml("test.yaml", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	os.Setenv("USER_NAME", "thanos")
	os.Setenv("USER_ADDRESS_STREET", "newstreet")
	os.Setenv("USER_AGE", "999")
	os.Setenv("USER_TALK", "true")
	SetValuesFromEnvironmentTag(&user)
	if user.Name != "thanos" {
		t.Errorf("Name was not set from env var.")
	}

	if user.Address.Street != "newstreet" {
		t.Errorf("Street was not set in Address.")
	}

	if user.Age != 999 {
		t.Errorf("Age was not set from env var.")
	}

	if !user.CanTalk {
		t.Errorf("user cannot talk")
	}

	os.Setenv("USER_NAME", "")
	os.Setenv("USER_ADDRESS_STREET", "")
	os.Setenv("USER_AGE", "")
	os.Setenv("USER_TALK", "")

}

func TestSetEnvDynamic(t *testing.T) {
	var user User
	err := ParseYaml("test.yaml", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	os.Setenv("USER_AGE", "50")
	os.Setenv("USER_ADDRESS_CITY", "New Castle")
	os.Setenv("USER_LIMIT", "99")
	os.Setenv("USER_CANWALK", "true")
	os.Setenv("USER_SCORE", "5.4242")
	SetValuesFromEnvironment("USER", &user)
	if user.Age != 50 {
		t.Errorf("Age was not set from env var.")
	}

	if user.Address.City != "New Castle" {
		t.Errorf("City was not set in Address.")
	}

	if user.Limit != 99 {
		t.Errorf("Limit has not been set from env var.")
	}

	if !user.CanWalk {
		t.Errorf("User cannot walk")
	}

	os.Setenv("USER_AGE", "")
	os.Setenv("USER_ADDRESS_CITY", "")
	os.Setenv("USER_LIMIT", "")
	os.Setenv("USER_CANWALK", "")
	os.Setenv("USER_SCORE", "")

}

func TestSetEnvDynamicNoPrefix(t *testing.T) {
	var user User
	err := ParseYaml("test.yaml", &user)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	os.Setenv("AGE", "50")
	os.Setenv("ADDRESS_CITY", "New Castle")
	os.Setenv("USER_LIMIT", "99")
	SetValuesFromEnvironment("", &user)
	if user.Age != 50 {
		t.Errorf("Age was not set from env var.")
	}

	if user.Address.City != "New Castle" {
		t.Errorf("City was not set in Address.")
	}

	if user.Limit != 10 {
		t.Errorf("Limit has not been set from env var.")
	}

	os.Setenv("AGE", "")
	os.Setenv("ADDRESS_CITY", "")
	os.Setenv("USER_LIMIT", "")

}
