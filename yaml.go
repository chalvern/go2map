package go2map

import (
	"fmt"

	"github.com/chalvern/sugar"
	"gopkg.in/yaml.v2"
)

// Yaml2MapFromFile convert content to go's flat map from file.
func Yaml2MapFromFile(fileName string) FlatMap {
	contentByte := mustReadFile(fileName)
	return Yaml2Map(contentByte)
}

// Yaml2Map convert yaml to go's flat map
func Yaml2Map(yamlData []byte) FlatMap {
	rawMap := make(map[interface{}]interface{})
	err := yaml.Unmarshal(yamlData, &rawMap)
	if err != nil {
		sugar.Fatalf("error: %v", err)
	}

	targetMap := make(FlatMap)
	nestedMap2FlatMap(rawMap, targetMap, "")
	return targetMap
}

// nestedMap2FlatMap convert nested Map to flat map
// @rawMap nested map
// @targetMap target flat map
// @parentPath the parent hierarchy 'x.x' of 'x.x.x'
func nestedMap2FlatMap(rawMap map[interface{}]interface{}, targetMap FlatMap, parentPath string) {
	if parentPath != "" {
		parentPath += "."
	}
	for rawKey, rawValue := range rawMap {
		rawKey := fmt.Sprintf("%v", rawKey)
		switch v := rawValue.(type) {
		case map[interface{}]interface{}:
			nestedMap2FlatMap(v, targetMap, parentPath+rawKey)
		default:
			targetMap[parentPath+rawKey] = rawValue
		}
	}
}
