package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

func main() {
	data := []byte(`{
  		"person": {
    		"name": {
        		"first": "Leonid",
				"last": "Bugaev",
				"fullName": "Leonid Bugaev"
    		},
    		"github": {
      			"handle": "buger",
      			"followers": 109
    		},
    		"avatars": [
      			{ 
					"url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", 
					"type": "thumbnail" 
				}
    		]
  		},
  		"company": {
    		"name": "Acme"
		}
	}`)
	val, dataType, offset, err := jsonparser.Get(data, "person", "name", "fullName")
	fmt.Println(string(val), dataType, offset, err)

	val, dataType, offset, err = jsonparser.Get(data, "person", "github", "followers")
	fmt.Println(string(val), dataType, offset, err)
	/*// json number类型转int
	intSli := make([]byte, 0, 4)
	if len(val) < 4 {
		intSli = append(intSli, 0)
		intSli = append(intSli, val...)
	}
	numberToInt := binary.LittleEndian.Uint32(intSli)
	fmt.Println(numberToInt)*/

	jpInt, err := jsonparser.GetInt(data, "person", "github", "followers")
	fmt.Println(jpInt, err)

	jpInt, err = jsonparser.GetInt(data, "person", "github", "handle")
	fmt.Println(jpInt, err)

	github := struct {
		Handle    string `json:"handle"`
		Followers int    `json:"followers"`
	}{}
	err = jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		switch string(key) {
		case "handle":
			github.Handle = string(value)
		case "followers":
			followers, _ := jsonparser.ParseInt(value)
			github.Followers = int(followers)
		}
		return nil
	}, "person", "github")
	fmt.Printf("github = %+v\n\n", github)

	byteGitHub := []byte(fmt.Sprintf(`{"handle: %s", "followers": "%d"}`, github.Handle, github.Followers))
	fmt.Println(byteGitHub, string(byteGitHub))
	// jsonparser.Set()
	githubJson, err := jsonparser.Set([]byte("{}"), byteGitHub, "github")
	fmt.Println(err, string(githubJson))

	paths := [][]string{
		{"person", "name", "fullName"},
		{"person", "avatars", "[0]", "url"},
		{"company", "name"},
	}

	jsonparser.EachKey(data, func(i int, bytes []byte, valueType jsonparser.ValueType, err error) {
		switch i {
		case 0:
			fmt.Printf("fullName = %s\n", bytes)
		case 1:
			fmt.Printf("avatars[0].url = %s\n", bytes)
		case 2:
			fmt.Printf("company.name = %s\n\n", bytes)
		}
	}, paths...)

}
