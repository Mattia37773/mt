/*
Copyright © 2026 Matze
*/
package yaml

import (
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/mattia37773/mt/functions/ui"
)

func GetConfig(yamlValue string, defaultValue string) string {
	yml, err := os.ReadFile(".mt.yaml")
	if err != nil {
		return defaultValue
	}

	var data interface{}
	if err := yaml.Unmarshal(yml, &data); err != nil {
		return defaultValue
	}

	parts := strings.Split(yamlValue, ".")

	results := findByPath(data, parts)

	if len(results) == 1 {
		return results[0]
	}

	if len(results) > 1 {
		fmt.Println(ui.Red("ERROR: multiple values found for"), yamlValue)
		return defaultValue
	}

	return defaultValue
}

func findByPath(data interface{}, path []string) []string {
	if len(path) == 0 {
		return []string{fmt.Sprintf("%v", data)}
	}

	current := path[0]

	if strings.Contains(current, "[*]") {
		key := strings.ReplaceAll(current, "[*]", "")

		switch t := data.(type) {
		case map[string]interface{}:
			if arr, ok := t[key].([]interface{}); ok {
				var results []string
				for _, item := range arr {
					results = append(results, findByPath(item, path[1:])...)
				}
				return results
			}
		}
	}

	switch t := data.(type) {
	case map[string]interface{}:
		if val, ok := t[current]; ok {
			return findByPath(val, path[1:])
		}
	case map[interface{}]interface{}:
		for k, v := range t {
			if fmt.Sprintf("%v", k) == current {
				return findByPath(v, path[1:])
			}
		}
	}

	return nil
}
