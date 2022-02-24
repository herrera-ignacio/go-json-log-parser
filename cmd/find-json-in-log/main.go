package main

import (
	"fmt"
	parser "logparser"
	"os"
	//"reflect"
)

//func compareJsonStringDeepEqual(j1, j2 []byte) bool {
//	var jToFind interface{}
//	err := json.Unmarshal(j1, &jToFind)
//	parser.Check(err)
//
//	var jToCompare interface{}
//	err = json.Unmarshal(j2, &jToCompare)
//	parser.Check(err)
//
//	return reflect.DeepEqual(jToFind, jToCompare)
//}

func main() {
	var found bool
	jsonRaw, err := os.ReadFile("./input.json")
	parser.Check(err)

	text, err := os.ReadFile("./log.txt")
	parser.Check(err)

	jsons := parser.GetAllJSONStrings(string(text))

	for i := 0; i < len(jsons) && !found; i++ {
		fmt.Println(jsons[i])
		if string(jsonRaw) == jsons[i] {
			found = true
		}
	}

	fmt.Printf("found: %t", found)
}
