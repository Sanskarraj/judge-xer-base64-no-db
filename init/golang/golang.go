package golang

import (
	// "encoding/json"
	"fmt"
	"strings"
)


func processValue (v interface{}) string {
	switch val := v.(type) {
	case []interface{}:
		items := make([]string, len(val))
		for i, item := range val {
			items[i] = processValue(item)
		}
		return fmt.Sprintf("{%s}", strings.Join(items, ", "))
	case map[string]interface{}:
		// Handle objects (not needed in this case, but included for completeness)
		items := make([]string, 0, len(val))
		for _, v := range val {
			items = append(items, processValue(v))
		}
		return fmt.Sprintf("{%s}", strings.Join(items, ", "))
	case float64:
		return fmt.Sprintf("%d", int(val))
	case string:
		return fmt.Sprintf("%q", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func GetBaseType (input interface{}) string {
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

func GetDepth (input interface{}) int {
	if array, ok := input.([]interface{}); ok && len(array) > 0 {
		return 1 + GetDepth(array[0])
	}
	return 0
}

func GetNesting(depth int) string {
  return strings.Repeat("[]", depth)
}



func ProcessJSON (jsonData interface{}, name string ) string {

    // var data interface{}
	// err := json.Unmarshal([]byte(jsonData), &data)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	result := processValue(jsonData)
  depth:=GetDepth(jsonData)
	
nesting:= GetNesting(depth)

baseType:= GetBaseType(jsonData)
name = name+"_s"
	s:= name+" := "+nesting+baseType+result
  return s;
  }