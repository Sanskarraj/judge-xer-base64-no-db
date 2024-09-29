package javascript

import (
	"encoding/json"
	// "fmt"
	//  "strings"
)




  func ProcessJSON (jsonData interface{}, name string) string {


	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return ""
	}

	jsonString := string(jsonBytes)
	
	name = name+"_s"

	s:="const "+ name+" = "+jsonString
  return s;
  }


