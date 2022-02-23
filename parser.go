// Package parser: given a string return a list of JSON objects.
package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// Parse a string as JSON if valid format or return nil.
func parseJson(text string) interface{} {
	var res interface{}

	err := json.Unmarshal([]byte(text), &res)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return res
}

// Returns RegEx to identify valid JSON in a string.
// It will detect JSON with nested objects up to "depth" level.
func getJsonRegex(depth int) string {
	jsonRegexBase := "{(?:[^{}]|(?R))*}"
	jsonRegex := "{(?:[^{}]|(?R))*}"
	r := regexp.MustCompile("\\(\\?R\\)")

	for i := 0; i < depth; i++ {
		jsonRegex = r.ReplaceAllString(jsonRegex, "(?:"+jsonRegexBase+")")
	}

	return r.ReplaceAllString(jsonRegex, "")
}

// GetAllJSONStrings Search a valid JSON in a given string
func GetAllJSONStrings(text string) []string {
	r, _ := regexp.Compile(getJsonRegex(0))
	return r.FindAllString(text, -1)
}
