package dotnet

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ConvertJSONToCSharpList converts nested JSON arrays into a formatted C# string.
func ConvertJSONToCSharpList(input interface{}, depth int) string {
	var buffer bytes.Buffer

	switch value := input.(type) {
	case []interface{}:
		if len(value) == 0 {
			return "new List<object>()"
		}

		baseType := GetBaseType(value[0])
		switch depth {
		case 1:
			buffer.WriteString(fmt.Sprintf("new List<%s> {", baseType))
		case 2:
			buffer.WriteString(fmt.Sprintf("%snew List<List<%s>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 3:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<%s>>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 4:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<List<%s>>>> {\n", strings.Repeat("    ", depth-1), baseType))
		case 5:
			buffer.WriteString(fmt.Sprintf("%snew List<List<List<List<List<%s>>>>> {\n", strings.Repeat("    ", depth-1), baseType))
		default:
			buffer.WriteString(fmt.Sprintf("%snew List<object> {\n", strings.Repeat("    ", depth-1)))
		}

		for i, v := range value {
			if depth > 1 {
				buffer.WriteString(strings.Repeat("    ", depth))
			}
			buffer.WriteString(ConvertJSONToCSharpList(v, depth-1))
			if i < len(value)-1 {
				buffer.WriteString(",")
				if depth > 1 {
					buffer.WriteString("\n")
				} else {
					buffer.WriteString(" ")
				}
			}
		}

		if depth > 1 {
			buffer.WriteString("\n" + strings.Repeat("    ", depth-1))
		}
		buffer.WriteString("}")

	case float64:
		buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
	case string:
		buffer.WriteString(fmt.Sprintf("\"%s\"", value))
	}
	return buffer.String()
}

// GetBaseType determines the base type of the nested array (int or string).
func GetBaseType(input interface{}) string {
	switch input.(type) {
	case []interface{}:
		if len(input.([]interface{})) > 0 {
			return GetBaseType(input.([]interface{})[0])
		}
		return "object"
	case float64:
		return "int"
	case string:
		return "string"
	default:
		return "object"
	}
}

// GenerateCSharpType generates the appropriate C# type based on the depth and base type of the nested array.
func GenerateCSharpType(depth int, baseType string) string {
	if depth == 0 {
		return baseType
	}
	return fmt.Sprintf("List<%s>", GenerateCSharpType(depth-1, baseType))
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}


func ProcessJSON(jsonData interface{},  name string ) string {

	// var jsonParsed interface{}

	// err := json.Unmarshal([]byte(jsonData), &jsonParsed)
	// if err != nil {
	// 	fmt.Println("Error parsing JSON:", err)
	// 	return err.Error()
	// }

	name = name+"_s"


	depth := GetDepth(jsonData)
	baseType := GetBaseType(jsonData)

	csharpType := GenerateCSharpType(depth, baseType)

	converted := ConvertJSONToCSharpList(jsonData, depth)

	// fmt.Printf("%s %s = %s;\n", csharpType, name, converted)


	s := csharpType + " " + name + " = " + converted + ";"
	return s
}






