package extract_test

import (
	"reflect"
	"sort"
	"testing"

	extract "github.com/meximonster/extractJsonKeys"
)

var o1 = []byte(`{
    "_id": "6207fa1a6ff619dd079e59e5",
    "index": 0,
    "guid": "354c1dc7-e517-45b3-af35-0c2389de7930",
    "isActive": false,
    "balance": "$1,246.33",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "blue",
    "name": "Mae Fuentes",
    "gender": "female",
    "company": "MOTOVATE",
    "email": "maefuentes@motovate.com",
    "phone": "+1 (926) 438-3720",
    "address": "785 Harman Street, Wright, Delaware, 9801",
    "about": "Irure duis reprehenderit duis veniam duis eiusmod culpa eiusmod esse Lorem elit. Lorem nulla officia ipsum culpa enim eiusmod esse eiusmod. Est adipisicing duis mollit incididunt exercitation sunt exercitation ex labore. Eiusmod Lorem incididunt elit Lorem enim eiusmod ad commodo. Tempor aliqua ea quis ex elit consectetur nulla adipisicing enim ullamco nostrud reprehenderit. Tempor ex ex exercitation quis labore consequat deserunt veniam eu esse culpa. Commodo nostrud non ea veniam voluptate deserunt ea in ipsum eiusmod est irure qui Lorem.\r\n",
    "registered": "2014-12-17T10:19:22 -02:00",
    "latitude": 43.83708,
    "longitude": 129.966523,
    "tags": [
      "nisi",
      "sunt",
      "fugiat",
      "consectetur",
      "occaecat",
      "excepteur",
      "non"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Olson Espinoza"
      },
      {
        "id": 1,
        "name": "Walker Saunders"
      },
      {
        "id": 2,
        "name": "Frost Ellis"
      }
    ],
    "greeting": "Hello, Mae Fuentes! You have 3 unread messages.",
    "favoriteFruit": "apple"
  }`)

var o2 = []byte(`{
	"random": 4,
	"random float": 6.35,
	"bool": true,
	"date": "1994-06-05",
	"regEx": "hello world",
	"enum": "online",
	"firstname": "Sharlene",
	"lastname": "Faro",
	"array": [
		"Mady",
		"Julieta",
		"Tera",
		"Phylis",
		"Bernardine"
	],
	"array of objects": [{
			"index": 0,
			"index start at 5": 5
		},
		{
			"index": 1,
			"index start at 5": 6
		}
	]
}`)

type testData struct {
	id     int
	input  []byte
	output []string
}

var data = []testData{{
	id:     1,
	input:  o1,
	output: []string{"_id", "about", "address", "age", "balance", "company", "email", "eyeColor", "favoriteFruit", "friends", "gender", "greeting", "guid", "id", "index", "isActive", "latitude", "longitude", "name", "phone", "picture", "registered", "tags"},
}, {
	id:     2,
	input:  o2,
	output: []string{"array", "array of objects", "bool", "date", "enum", "firstname", "index", "index start at 5", "lastname", "random", "random float", "regEx"},
},
}

func TestExtractJSONKeys(t *testing.T) {
	for i := range data {
		keys := extract.ExtractJSONKeys(data[i].input)
		sort.Strings(keys)
		if !reflect.DeepEqual(data[i].output, keys) {
			t.Errorf("slices for test %d not equal! Expected: %v, got %v", data[i].id, data[i].output, keys)
		}
	}
}

func TestExtractJSONValues(t *testing.T) {
	var output = []string{"1994-06-05", "Bernardine", "Faro", "Julieta", "Mady", "Phylis", "Sharlene", "Tera", "hello world", "online"}
	values := extract.ExtractJSONValues(o2)
	sort.Strings(values)
	equal := reflect.DeepEqual(output, values)
	if !equal {
		t.Errorf("slices not equal!")
	}
}
