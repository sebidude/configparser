[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/sebidude/configparser)](https://goreportcard.com/report/gitlab.com/sebidude/configparser)
# configparser
Parse a json or yaml file and unmarshal the content to a struct type

To set the values of the struct members from environment variables, you can define an "env" struct tag
which defines the name of the envvar to set the value from and call ```configparser.SetValuesFromEnvironmentTag```
on the already parsed struct or use ```configparser.SetValuesFromEnvironment``` to lookup the envvar by the name of the fields.

## Example
This is the config yaml
```
name: sebidude
age: 99
species: unknown
street: see address
address:
  street: homestreet
  city: donk city
limit: 10
```

Define a stuct type 
```
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
}
```

### Parse the configfile
```
...
var user User
configparser.ParseYaml("config.yaml",&user)
...
```

### Use environment variables

Set values from env var using struct tags
```
configparser.SetValuesFromEnvironmentTag(&user)
```
The env vars stated in the tags will be looked up and set if not empty

To set the values from env vars without defining a struct tag, use a matching name for the variable and set the values with ```configparser.SetValuesFromEnvironment("CONFIG",&user)```

Example  
```bash
export CONFIG_ADDRESS_CITY="Londinium"
```
```golang
configparser.SetValuesFromEnvironment("CONFIG",&user)
fmt.Println(user.Address.City)
```




