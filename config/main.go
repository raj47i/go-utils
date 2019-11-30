package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

// StoreType represent a support configuration store type to use
type StoreType uint

// Allowed values for type StoreType
const (
	StoreENV = iota // 0
	StoreYamlFile
	StoreJSONFile
)

// ConfigStoreTypeText stores the textual representation of valid StoreType vakues
var ConfigStoreTypeText = [...]string{
	"env-vars",
	"yamml-file",
	"json-file",
}

// Returns string/text value of a StoreType
func (s StoreType) String() string {
	if s < StoreENV || s > StoreJSONFile {
		return "Unknown"
	}
	return ConfigStoreTypeText[s]
}

// LoadFromENV loads configuration values to a struct from Host's ENV vars,
// values are loaded based on the field tags
func LoadFromENV(configuration interface{}) error {
	t := reflect.ValueOf(configuration).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Type().Field(i)
		value := os.Getenv(field.Tag.Get("env"))
		if value == "" {
			value = field.Tag.Get("default")
		}

		switch field.Type.Name() {
		case "string":
			t.Field(i).SetString(value)
		case "int":
			v, _ := strconv.ParseInt(value, 10, 0)
			t.Field(i).SetInt(v)
		case "uint":
			v, _ := strconv.ParseUint(value, 10, 0)
			t.Field(i).SetUint(v)
		case "int64":
			v, _ := strconv.ParseInt(value, 10, 0)
			t.Field(i).SetInt(v)
		case "bool":
			v, _ := strconv.ParseBool(value)
			t.Field(i).SetBool(v)
		default:
			return fmt.Errorf("Unrecognized data type for configuration")
		}
	}
	return nil
}

// LoadFromJSONFile loads configuration values to a struct from Host's ENV vars,
// values are loaded based on the field tags
func LoadFromJSONFile(configuration interface{}, jsonFile string) error {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("cannot read %s: %v", jsonFile, err)
	}
	return json.Unmarshal(file, &configuration)
}
