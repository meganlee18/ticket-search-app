package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func DisplayAllResultFields(path string) string {
	var result []map[string]interface{}

	data := readFile(path)
	unmarshalledResult, err := unmarshalData(data, result)
	if err != nil {
		fmt.Println("error: ", err)
	}
	sortedFields := displaySortedFields(unmarshalledResult)

	return strings.Join(removeDuplicateValues(sortedFields),  "\n")
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func unmarshalData(data []byte, result []map[string]interface{}) ([]map[string]interface{}, error) {
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return result, err
}

func displaySortedFields(result []map[string]interface{}) []string {
	var fields []string

	for _, v := range result {
		for k := range v {
			fields = append(fields, k)
		}
	}

	sort.Strings(fields)
	return fields
}

func removeDuplicateValues(fields []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range fields {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}
