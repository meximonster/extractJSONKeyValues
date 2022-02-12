package main

import (
	"encoding/json"
	"fmt"
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

var o2 = []byte(`[
  {
    "_id": "62081f4d4ffd5d4a3f3aba15",
    "index": 0,
    "guid": "c169577e-db4e-4eeb-9287-f88cd81f2832",
    "isActive": false,
    "balance": "$1,611.80",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "brown",
    "name": "Liza Richmond",
    "gender": "female",
    "company": "UNISURE",
    "email": "lizarichmond@unisure.com",
    "phone": "+1 (891) 538-2473",
    "address": "347 Apollo Street, Homeland, Wisconsin, 1231",
    "about": "Nostrud minim consectetur consectetur voluptate reprehenderit occaecat occaecat sit deserunt ut. Ullamco occaecat est ex ullamco proident minim ipsum excepteur sint quis pariatur incididunt velit. Nostrud nulla esse incididunt ut culpa nostrud cillum non enim Lorem fugiat reprehenderit irure. Consequat dolore exercitation amet cupidatat. Consectetur ex nisi culpa Lorem eiusmod minim mollit dolor aliquip irure eiusmod ex do nostrud. Cupidatat fugiat mollit sit nulla cillum aute eiusmod deserunt do eu.\r\n",
    "registered": "2014-08-18T12:27:58 -03:00",
    "latitude": -71.185906,
    "longitude": 170.439489,
    "tags": [
      "quis",
      "ut",
      "adipisicing",
      "nostrud",
      "et",
      "proident",
      "nulla"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Nora Rios"
      },
      {
        "id": 1,
        "name": "Mcdonald Sweeney"
      },
      {
        "id": 2,
        "name": "Sherri Morton"
      }
    ],
    "greeting": "Hello, Liza Richmond! You have 1 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "62081f4da4e351e66ef61eb7",
    "index": 1,
    "guid": "1fe0dcab-8393-4204-96bd-22ce2269a16a",
    "isActive": true,
    "balance": "$1,731.78",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "blue",
    "name": "Ferguson Patrick",
    "gender": "male",
    "company": "ZENCO",
    "email": "fergusonpatrick@zenco.com",
    "phone": "+1 (942) 423-2898",
    "address": "674 Gotham Avenue, Bellamy, Guam, 1907",
    "about": "Incididunt cupidatat exercitation esse excepteur est commodo id elit dolor irure. Eiusmod veniam adipisicing aliquip dolor non ullamco commodo labore ex. Eiusmod esse aliquip amet aute ut amet dolore sunt ea occaecat. Reprehenderit eu aliquip cillum velit sint incididunt esse excepteur. Irure eiusmod in non ut labore elit do quis sint qui do consequat veniam labore.\r\n",
    "registered": "2015-11-16T04:26:37 -02:00",
    "latitude": 24.589849,
    "longitude": -126.645983,
    "tags": [
      "mollit",
      "ad",
      "commodo",
      "ipsum",
      "eu",
      "dolore",
      "culpa"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Wilson Little"
      },
      {
        "id": 1,
        "name": "Georgina Davidson"
      },
      {
        "id": 2,
        "name": "Davidson Huffman"
      }
    ],
    "greeting": "Hello, Ferguson Patrick! You have 10 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "62081f4d7808807a92b7ef32",
    "index": 2,
    "guid": "34af7275-18b2-4e48-89f5-6434a6b04d54",
    "isActive": true,
    "balance": "$2,614.31",
    "picture": "http://placehold.it/32x32",
    "age": 25,
    "eyeColor": "brown",
    "name": "Schroeder Merritt",
    "gender": "male",
    "company": "IMKAN",
    "email": "schroedermerritt@imkan.com",
    "phone": "+1 (803) 500-3328",
    "address": "591 Village Road, Ezel, North Carolina, 2543",
    "about": "Excepteur est nostrud sunt consectetur. Anim adipisicing eiusmod esse voluptate qui do in non consequat quis ea cupidatat. Sint sit magna ipsum voluptate enim qui reprehenderit nisi mollit officia aliquip anim anim non. Reprehenderit aliquip id amet cupidatat labore aliquip eu.\r\n",
    "registered": "2019-07-28T04:55:38 -03:00",
    "latitude": 78.635441,
    "longitude": -76.255116,
    "tags": [
      "quis",
      "duis",
      "proident",
      "occaecat",
      "aute",
      "quis",
      "cupidatat"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Rice Hendrix"
      },
      {
        "id": 1,
        "name": "Acosta Mckee"
      },
      {
        "id": 2,
        "name": "Forbes Austin"
      }
    ],
    "greeting": "Hello, Schroeder Merritt! You have 8 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "62081f4d685aa9de48b75982",
    "index": 3,
    "guid": "d93db587-d706-4e84-a0e6-fdd5b8a9e8be",
    "isActive": true,
    "balance": "$2,069.44",
    "picture": "http://placehold.it/32x32",
    "age": 40,
    "eyeColor": "green",
    "name": "Torres Mercer",
    "gender": "male",
    "company": "FILODYNE",
    "email": "torresmercer@filodyne.com",
    "phone": "+1 (832) 538-3020",
    "address": "718 Court Street, Basye, Northern Mariana Islands, 3279",
    "about": "Laborum officia enim ullamco pariatur id anim. Commodo dolore cupidatat laboris incididunt ea. Nulla eu laborum ut consectetur dolore sunt fugiat elit aliquip.\r\n",
    "registered": "2021-09-07T12:18:00 -03:00",
    "latitude": 54.140372,
    "longitude": -171.342851,
    "tags": [
      "nisi",
      "incididunt",
      "Lorem",
      "non",
      "reprehenderit",
      "eu",
      "tempor"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Durham Schultz"
      },
      {
        "id": 1,
        "name": "Ava Hammond"
      },
      {
        "id": 2,
        "name": "Nina Rogers"
      }
    ],
    "greeting": "Hello, Torres Mercer! You have 9 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "62081f4db0ca77cb8c4ddc2f",
    "index": 4,
    "guid": "63a6fed3-c51c-4bd4-95ae-d2155a2b495f",
    "isActive": true,
    "balance": "$1,419.72",
    "picture": "http://placehold.it/32x32",
    "age": 37,
    "eyeColor": "blue",
    "name": "Carolina Deleon",
    "gender": "female",
    "company": "DANCITY",
    "email": "carolinadeleon@dancity.com",
    "phone": "+1 (904) 554-3937",
    "address": "469 Borinquen Pl, Rodanthe, Maryland, 4774",
    "about": "Non laborum exercitation occaecat enim et cillum labore laboris. Et dolor in Lorem sint nostrud. Ex enim esse velit duis cupidatat adipisicing. Laboris in velit officia reprehenderit veniam reprehenderit non minim. Fugiat sint amet exercitation adipisicing in commodo enim reprehenderit esse Lorem Lorem ullamco. Aliqua mollit id minim cillum pariatur id sint.\r\n",
    "registered": "2015-11-09T05:12:13 -02:00",
    "latitude": 7.903213,
    "longitude": -169.489617,
    "tags": [
      "elit",
      "cupidatat",
      "proident",
      "anim",
      "do",
      "fugiat",
      "Lorem"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Steele Monroe"
      },
      {
        "id": 1,
        "name": "Foley Robertson"
      },
      {
        "id": 2,
        "name": "Mallory Crosby"
      }
    ],
    "greeting": "Hello, Carolina Deleon! You have 10 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "62081f4d65910c29b7b4f96a",
    "index": 5,
    "guid": "6448d9fe-753a-486c-a496-2c589336e9a9",
    "isActive": true,
    "balance": "$1,522.25",
    "picture": "http://placehold.it/32x32",
    "age": 26,
    "eyeColor": "blue",
    "name": "Amparo Mckinney",
    "gender": "female",
    "company": "GEEKKO",
    "email": "amparomckinney@geekko.com",
    "phone": "+1 (990) 564-3379",
    "address": "701 Lexington Avenue, Sperryville, California, 7166",
    "about": "Esse magna ex incididunt sit veniam mollit est sint minim duis aute magna. Excepteur eu quis et reprehenderit sit irure sint dolore reprehenderit proident fugiat. Fugiat mollit mollit id deserunt magna aliqua Lorem incididunt. Irure excepteur magna sint ex commodo veniam voluptate Lorem ad ipsum. Eu fugiat culpa anim dolore tempor incididunt voluptate eiusmod magna minim. Anim non consectetur in eiusmod adipisicing.\r\n",
    "registered": "2020-05-31T06:39:29 -03:00",
    "latitude": 35.075987,
    "longitude": -28.429611,
    "tags": [
      "ea",
      "laboris",
      "Lorem",
      "tempor",
      "aute",
      "anim",
      "elit"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Dianna Witt"
      },
      {
        "id": 1,
        "name": "Robyn Griffith"
      },
      {
        "id": 2,
        "name": "Sharron Zamora"
      }
    ],
    "greeting": "Hello, Amparo Mckinney! You have 4 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "62081f4d58f6c3d3e98a7f28",
    "index": 6,
    "guid": "46b79fa6-c79b-4fb3-a372-c6f0bb134bee",
    "isActive": true,
    "balance": "$1,086.93",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "brown",
    "name": "Leola Jacobs",
    "gender": "female",
    "company": "NORSUL",
    "email": "leolajacobs@norsul.com",
    "phone": "+1 (929) 554-2480",
    "address": "104 Batchelder Street, Spelter, Nebraska, 5000",
    "about": "Enim excepteur esse qui adipisicing dolore laborum irure laboris elit sunt anim. Sint pariatur Lorem exercitation qui do ea non fugiat velit tempor ad. Nulla enim consequat esse exercitation aliqua id. Nisi non do culpa aliqua. Anim labore nisi quis deserunt proident id excepteur eu elit ullamco id sunt veniam nisi. Incididunt incididunt dolore commodo incididunt occaecat.\r\n",
    "registered": "2020-02-15T04:08:09 -02:00",
    "latitude": -4.139411,
    "longitude": 116.56453,
    "tags": [
      "amet",
      "aute",
      "occaecat",
      "anim",
      "ex",
      "ipsum",
      "tempor"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Martin Oconnor"
      },
      {
        "id": 1,
        "name": "Fields Roth"
      },
      {
        "id": 2,
        "name": "Gray Stewart"
      }
    ],
    "greeting": "Hello, Leola Jacobs! You have 5 unread messages.",
    "favoriteFruit": "apple"
  }
]`)

var inputs = [][]byte{o1, o2}

func main() {
	var keys *[]string
	for i := range inputs {
		i := unmarshalJSON(inputs[i])
		switch t := i.(type) {
		case map[string]interface{}:
			keys = extractKeysFromJson(&[]string{}, t)
		case []interface{}:
			keys = extractKeysFromJsonSlice(&[]string{}, t)
		}
		uniqueKeys := removeDuplicates(*keys)
		fmt.Println(uniqueKeys)
	}
}

func unmarshalJSON(o []byte) interface{} {
	var m map[string]interface{}
	err := json.Unmarshal(o, &m)
	if err != nil {
		var s []interface{}
		err = json.Unmarshal(o, &s)
		return s
	}
	return m
}

func extractKeysFromJson(keys *[]string, m map[string]interface{}) *[]string {
	for key, value := range m {
		*keys = append(*keys, key)
		switch t := value.(type) {
		case map[string]interface{}:
			extractKeysFromJson(keys, t)
		case []interface{}:
			for item := range t {
				switch t := t[item].(type) {
				case map[string]interface{}:
					extractKeysFromJson(keys, t)
				}
			}
		}
	}
	return keys
}

func extractKeysFromJsonSlice(keys *[]string, s []interface{}) *[]string {
	for _, item := range s {
		switch t := item.(type) {
		case map[string]interface{}:
			extractKeysFromJson(keys, t)
		case []interface{}:
			for item := range t {
				switch t := t[item].(type) {
				case map[string]interface{}:
					extractKeysFromJson(keys, t)
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
