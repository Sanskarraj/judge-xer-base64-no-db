package typescript

import (
    "bytes"
    // "encoding/json"
    "fmt"
    "strconv"
    "strings"
)

// ConvertJSONToTypeScript converts JSON data to TypeScript array format.
func ConvertJSONToTypeScript(input interface{}, depth int) string {
    var buffer bytes.Buffer
    switch value := input.(type) {
    case []interface{}:
        if depth == 1 {
            buffer.WriteString("[")
            for i, v := range value {
                buffer.WriteString(ConvertJSONToTypeScript(v, depth-1))
                if i < len(value)-1 {
                    buffer.WriteString(", ")
                }
            }
            buffer.WriteString("]")
        } else {
            buffer.WriteString("[\n")
            for i, v := range value {
                buffer.WriteString(strings.Repeat("  ", depth))
                buffer.WriteString(ConvertJSONToTypeScript(v, depth-1))
                if i < len(value)-1 {
                    buffer.WriteString(",\n")
                }
            }
            buffer.WriteString("\n" + strings.Repeat("  ", depth-1) + "]")
        }
    case float64:
        buffer.WriteString(strconv.FormatFloat(value, 'f', -1, 64))
    case string:
        buffer.WriteString(fmt.Sprintf("\"%s\"", value))
    }
    return buffer.String()
}

func GetNesting(depth int) string {
    return strings.Repeat("[]", depth)
}

// GetDepth determines the depth of the nested array.
func GetDepth(input interface{}) int {
    if array, ok := input.([]interface{}); ok && len(array) > 0 {
        return 1 + GetDepth(array[0])
    }
    return 0
}

// GetType determines the TypeScript type of the input.
func GetType(input interface{}) string {
    switch input.(type) {
    case []interface{}:
        if len(input.([]interface{})) > 0 {
            return GetType(input.([]interface{})[0])
        }
        return "any"
    case float64:
        return "number"
    case string:
        return "string"
    default:
        return "any"
    }
}


func ProcessJSON(jsonData interface{},  name string ) string {

    // var jsonParsed interface{}
    // err := json.Unmarshal([]byte(jsonData), &jsonParsed)
    // if err != nil {
    //     return fmt.Sprintf("Error parsing JSON: %s", err)
    // }

    depth := GetDepth(jsonData)
    converted := ConvertJSONToTypeScript(jsonData, depth)
    typeStr := GetType(jsonData)
    
	name = name+"_s"


    return fmt.Sprintf("const %s: %s%s = %s; ", name, typeStr, GetNesting(depth), converted)
}

