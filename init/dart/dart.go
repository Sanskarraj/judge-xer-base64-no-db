package dart

import (
	// "fmt"
  "encoding/json"
)




func ProcessJSON(jsonData interface{},  name string ) string  {
	name = name+"_s"


	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return ""
	}

	jsonString := string(jsonBytes)
	

	return "var " + name+" = "+jsonString + ";"
}


