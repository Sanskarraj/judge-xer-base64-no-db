
package cpp

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	// "log"
	"strings"
)






// formatToCppVector converts a JSON structure to a C++ vector string representation
func formatToCppVector(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		elements := make([]string, len(v))
		for i, elem := range v {
			elements[i] = formatToCppVector(elem)
		}
		return fmt.Sprintf("{%s}", strings.Join(elements, ", "))
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case float64:
		if v == float64(int(v)) {
			return fmt.Sprintf("%d", int(v))
		}
		return fmt.Sprintf("%g", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// getVectorType determines the C++ vector type based on the content
func getVectorType(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		if len(v) == 0 {
			return "auto"
		}
		innerType := getVectorType(v[0])
		if strings.HasPrefix(innerType, "std::vector") {
			return fmt.Sprintf("std::vector<%s>", innerType)
		}
		return fmt.Sprintf("std::vector<%s>", innerType)
	case string:
		return "std::string"
	case float64:
		return "int"
	}
	return "auto"
}

// parseAndCleanJSON removes outer curly braces and parses the JSON input
// func parseAndCleanJSON(jsonInput string) (interface{}, error) {
// 	// Remove the outer curly braces and any whitespace
// 	jsonInput = strings.TrimSpace(jsonInput)
// 	jsonInput = strings.TrimPrefix(jsonInput, "{")
// 	jsonInput = strings.TrimSuffix(jsonInput, "}")
// 	jsonInput = strings.TrimSpace(jsonInput)

// 	var data interface{}
// 	if err := json.Unmarshal([]byte(jsonInput), &data); err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }

// converterCplusplus converts a JSON string to a C++ vector representation
func ProcessJSON(jsonData interface{},  name string ) string {
	// data, err := parseAndCleanJSON(input)
	// if err != nil {
	// 	fmt.Printf("Error decoding JSON: %v\n", err)
	// 	return ""
	// }

	name = name+"_s"

	vectorType := getVectorType(jsonData)
	result := fmt.Sprintf("%s%s%s%s", vectorType,name,formatToCppVector(jsonData),";")
	return result
}
