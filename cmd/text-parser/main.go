package main

import (
	"fmt"
	parser "logparser"
	"os"
)

func main() {
	data, err := os.ReadFile("./log.txt")

	if err != nil {
		fmt.Println("You need to provide an 'log.txt' file")
	}

	fmt.Println("JSONs:")

	jsons := parser.GetAllJSONStrings(string(data))

	if len(jsons) == 0 {
		fmt.Println("No JSON found!")
		os.Exit(1)
	}

	for i, json := range jsons {
		fmt.Printf("##### %d #####\n", i+1)
		fmt.Println(json)
	}

	fmt.Println("Done!")
}
