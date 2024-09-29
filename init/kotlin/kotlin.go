package kotlin

import (
	"bytes"
	"fmt"
	"strings"
)

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("\"%v\"", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func parseArray(array []interface{}) string {
	var buffer bytes.Buffer

	// Use mutableListOf instead of listOf
	buffer.WriteString("mutableListOf(")
	for i, element := range array {
		switch elem := element.(type) {
		case []interface{}:
			buffer.WriteString(parseArray(elem))
		default:
			buffer.WriteString(formatValue(elem))
		}
		if i != len(array)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString(")")
	return buffer.String()
}

func parseJSON(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		return parseArray(v)
	default:
		return formatValue(v)
	}
}

func ProcessJSON(jsonData interface{}, name string) string {
	name = name + "_s"

	// Replace listOf with mutableListOf
	result := "val " + name + " = " + parseJSON(jsonData)

	// Properly format the final output for display
	s := strings.ReplaceAll(result, "mutableListOf(", "\nmutableListOf(")
	return s
}
