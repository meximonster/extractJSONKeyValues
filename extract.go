package extract

import (
	"encoding/json"
	"log"
)

func ExtractJSONKeys(o []byte) []string {
	var keys *[]string
	i := unmarshalJSON(o)
	switch t := i.(type) {
	case map[string]interface{}:
		keys = GetKeysFromJson(&[]string{}, t)
	case []interface{}:
		keys = GetKeysFromJsonSlice(&[]string{}, t)
	}
	uniqueKeys := removeDuplicates(*keys)
	return uniqueKeys
}

func unmarshalJSON(o []byte) interface{} {
	var m map[string]interface{}
	err := json.Unmarshal(o, &m)
	if err != nil {
		var s []interface{}
		err = json.Unmarshal(o, &s)
		if err != nil {
			log.Fatal(err)
		}
		return s
	}
	return m
}

func GetKeysFromJson(keys *[]string, m map[string]interface{}) *[]string {
	for key, value := range m {
		*keys = append(*keys, key)
		switch t := value.(type) {
		case map[string]interface{}:
			GetKeysFromJson(keys, t)
		case []interface{}:
			for item := range t {
				switch t := t[item].(type) {
				case map[string]interface{}:
					GetKeysFromJson(keys, t)
				}
			}
		}
	}
	return keys
}

func GetKeysFromJsonSlice(keys *[]string, s []interface{}) *[]string {
	for _, item := range s {
		switch t := item.(type) {
		case map[string]interface{}:
			GetKeysFromJson(keys, t)
		case []interface{}:
			for item := range t {
				switch t := t[item].(type) {
				case map[string]interface{}:
					GetKeysFromJson(keys, t)
				}
			}
		}
	}
	return keys
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
