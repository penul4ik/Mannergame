package main

import (
	"fmt"
	"os"

	"github.com/clbanning/mxj/v2"
	"gopkg.in/yaml.v2"
)

type Map map[string]any

func xmlDecode(pathToFile string) (map[string]any, error) {
	byteValue, err := os.ReadFile(pathToFile)
	if err != nil {
		fmt.Printf("Error while read file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Seccessfully read %v\n", pathToFile)
	mv, err := mxj.NewMapXml(byteValue)
	if err != nil {
		fmt.Printf("Error convert xml to map: %v\n", err)
		os.Exit(1)
	}
	return mv, nil
}

func cleanMap(m map[string]any) map[string]any {
	clean := make(map[string]any)
	for k, v := range m {
		key := k
		if len(k) > 1 && k[0] == '-' {
			key = k[1:]
		}

		// рекурсивно чистим вложенные мапы
		switch val := v.(type) {
		case map[string]any:
			clean[key] = cleanMap(val)
		case []any:
			newList := make([]any, len(val))
			for i, item := range val {
				if submap, ok := item.(map[string]any); ok {
					newList[i] = cleanMap(submap)
				} else {
					newList[i] = item
				}
			}
			clean[key] = newList
		default:
			clean[key] = val
		}
	}
	return clean
}

func args() []string {
	file := os.Args
	if len(file) < 2 {
		err := fmt.Errorf("specify the path to the XML file")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return file
}

func main() {
	path := args()
	data, err := xmlDecode(string(path[1]))
	if err != nil {
		fmt.Printf("Error while decode xml: %v", err)
		os.Exit(1)
	}
	dataCleaned := cleanMap(data)
	yamlData, err := yaml.Marshal(&dataCleaned)
	if err != nil {
		fmt.Printf("Error while marshaling: %v", err)
		os.Exit(1)
	}
	fmt.Println(string(yamlData))
	os.Exit(0)
}
