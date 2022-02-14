package extract

import (
	"encoding/json"
	"log"
)

func ExtractJSONKeys(o []byte) []string {
	var keys *[]string
	var m map[string]interface{}
	err := json.Unmarshal(o, &m)
	if err != nil {
		log.Fatal(err)
	}
	keys = getKeysFromJson(&[]string{}, m)
	uniqueKeys := removeDuplicates(*keys)
	return uniqueKeys
}

func ExtractJSONValues(o []byte) []string {
	var values *[]string
	var m map[string]interface{}
	err := json.Unmarshal(o, &m)
	if err != nil {
		log.Fatal(err)
	}
	values = getValuesFromJson(&[]string{}, m)
	uniqueKeys := removeDuplicates(*values)
	return uniqueKeys
}

func getKeysFromJson(keys *[]string, m map[string]interface{}) *[]string {
	for key, value := range m {
		*keys = append(*keys, key)
		switch t := value.(type) {
		case map[string]interface{}:
			getKeysFromJson(keys, t)
		case []interface{}:
			for item := range t {
				switch t := t[item].(type) {
				case map[string]interface{}:
					getKeysFromJson(keys, t)
				}
			}
		}
	}
	return keys
}

func getValuesFromJson(values *[]string, m map[string]interface{}) *[]string {
	for _, value := range m {
		switch t := value.(type) {
		case string:
			*values = append(*values, t)
		case map[string]interface{}:
			getValuesFromJson(values, t)
		case []interface{}:
			for i := range t {
				switch typ := t[i].(type) {
				case string:
					*values = append(*values, typ)
				case map[string]interface{}:
					getValuesFromJson(values, typ)
				}
			}
		}
	}
	return values
}

func removeDuplicates(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
