package java

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ConvertJSONToJavaList converts nested JSON arrays into a formatted Java-like string.
func ConvertJSONToJavaList(input interface{}, depth int) string {
	var buffer bytes.Buffer

	switch value := input.(type) {
	case []interface{}:
		if depth == 1 {
			buffer.WriteString("Arrays.asList(")
			for i, v := range value {
				buffer.WriteString(ConvertJSONToJavaList(v, depth-1))
				if i < len(value)-1 {
					buffer.WriteString(", ")
				}
			}
			buffer.WriteString(")")
		} else {
			buffer.WriteString(strings.Repeat("    ", depth-1) + "Arrays.asList(\n")
			for i, v := range value {
				buffer.WriteString(strings.Repeat("    ", depth))
				buffer.WriteString(ConvertJSONToJavaList(v, depth-1))
				if i < len(value)-1 {
					buffer.WriteString(",\n")
				}
			}
			buffer.WriteString("\n" + strings.Repeat("    ", depth-1) + ")")
		}
	case float64:
		buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	case string:
		buffer.WriteString(fmt.Sprintf("\"%s\"", value))
	}
	return buffer.String()
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}

// GetBaseType determines the base type of the nested array (Integer or String).
func GetBaseType(input interface{}) string {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return GetBaseType(array[0])
	}
	switch input.(type) {
	case float64:
		return "Integer"
	case string:
		return "String"
	default:
		return "Object"
	}
}

// GenerateJavaType generates the appropriate Java type based on the depth and base type of the nested array.
func GenerateJavaType(depth int, baseType string) string {
	if depth == 0 {
		return baseType
	}
	return fmt.Sprintf("List<%s>", GenerateJavaType(depth-1, baseType))
}



func ProcessJSON(jsonData interface{},  name string ) string {
	// var jsonParsed interface{}

	// // Parse the JSON string into a generic interface{}
	// err := json.Unmarshal([]byte(jsonData), &jsonParsed)
	// if err != nil {
	// 	fmt.Println("Error parsing JSON:", err)
	// 	return err.Error() // Return the error message as a string
	// }

	// // Determine the name based on the input parameter 'in'
	// name := "nums"
	// if in == 1 {
	// 	name = "targets"
	// }

	// Determine the depth and base type of the parsed JSON structure
	depth := GetDepth(jsonData)
	baseType := GetBaseType(jsonData)

	// Generate the corresponding Java type (e.g., List<List<Integer>>)
	javaType := GenerateJavaType(depth, baseType)

	// Convert the parsed JSON into a Java-style list declaration
	converted := ConvertJSONToJavaList(jsonData, depth)

	// Print the Java variable declaration (optional)
	//fmt.Printf("%s %s = %s;\n", javaType, name, converted)

	// Return the Java variable declaration as a string
	name = name+"_s"
	s := javaType + " " + name + " = " + converted + ";"
	return s
}





