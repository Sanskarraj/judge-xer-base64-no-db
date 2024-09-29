package rust

import (
	"bytes"
	// "encoding/json"
	"fmt"
)

// Format JSON data into Rust-style vec![] string format
func formatToVecString(data interface{}) string {
	switch v := data.(type) {
	case []interface{}:
		var buffer bytes.Buffer
		buffer.WriteString("vec![")
		for i, item := range v {
			buffer.WriteString(formatToVecString(item))
			if i < len(v)-1 {
				buffer.WriteString(", ")
			}
		}
		buffer.WriteString("]")
		return buffer.String()
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("\"%v\"", v)
	default:
		return ""
	}
}

// Process JSON and convert it to Rust variable assignments
func ProcessJSON(jsonData interface{},  name string ) string {

	
	result := formatToVecString(jsonData)

	name = name+"_s"
	

	return fmt.Sprintf("let %s = %s;", name, result)
}

