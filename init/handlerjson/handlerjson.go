package handlerjson


import (
	"encoding/json"
	"encoding/base64"
	"fmt"
	// "strings"
	// "back/golang"
)



// func main() {

// 	base64JSON:=`ewoJCSJpbnB1dCI6IHsKCQkJImlucHV0cGFyYTEiOiBbWzEsIDIsIDNdLCBbNCwgNSwgNl1dLAoJCQkiaW5wdXRwYXJhMiI6IFs3LCA4LCA5XSwKCQkJImlucHV0cGFyYTMiOiBbW1sxMCwgMTFdLCBbMTIsIDEzXV0sIFtbMTQsIDE1XSwgWzE2LCAxN11dXSwKCQkJImlucHV0cGFyYTQiOiBbWzE4LCAxOSwgMjBdLCBbMjEsIDIyLCAyM11dLAoJCQkiaW5wdXRwYXJhNSI6IFsyNCwgMjUsIDI2XSwKCQkJImlucHV0cGFyYU4iOiBbMjcsIDI4LCAyOV0sCgkJCSJpbnB1dHBhcmFOMSI6IFs3LCA4XSwKCQkJImlucHV0cGFyYU4yIjogW1tbMTAsIDExXSwgWzEyLCAxM11dXQoJCX0sCgkJIm91dHB1dCI6IHsKCQkJIm91dCI6IFtbWzMwLCAzMV0sIFszMiwgMzNdXSwgW1szNCwgMzVdLCBbMzYsIDM3XV1dCgkJfQoJfQ==`
// 	jsonString := decode(base64JSON)



// 	// jsonString := `{
// 	// 	"input": {
// 	// 		"inputpara1": [[1, 2, 3], [4, 5, 6]],
// 	// 		"inputpara2": [7, 8, 9],
// 	// 		"inputpara3": [[[10, 11], [12, 13]], [[14, 15], [16, 17]]],
// 	// 		"inputpara4": [[18, 19, 20], [21, 22, 23]],
// 	// 		"inputpara5": [24, 25, 26],
// 	// 		"inputparaN": [27, 28, 29],
// 	// 		"inputparaN1": [7, 8],
// 	// 		"inputparaN2": [[[10, 11], [12, 13]]]
// 	// 	},
// 	// 	"output": {
// 	// 		"out": [[[30, 31], [32, 33]], [[34, 35], [36, 37]]]
// 	// 	}
// 	// }`

// 	// Extract and print data
// 	inputData := extractInputData(jsonString)
// 	outputData := extractOutputData(jsonString)

// 	// Print values in the order of keys from innerInputParameter
// 	fmt.Println("\nInput values in order:")
// 	for _, key := range innerInputParameter(jsonString) {
// 		if val, exists := inputData[key]; exists {
// 			fmt.Printf("%s: %v\n", key, val)
// 		}
// 	}

// 	// Print values in the order of keys from innerOutputParameter
// 	fmt.Println("\nOutput values in order:")
// 	for _, key := range innerOutputParameter(jsonString) {
// 		if val, exists := outputData[key]; exists {
// 			fmt.Printf("%s: %v\n", key, val)
// 		}
// 	}
// }



func decode (s string) string{
	decodedBytes,_ := base64.StdEncoding.DecodeString(s)
	// if err != nil {
	// 	fmt.Println("Error decoding:", err)
	// 	return
	// }
	decodedString := string(decodedBytes)
	// fmt.Println("Decoded (Original):", decodedString)

	return decodedString
}




func ExtractDataMap(base64JSON string) (map[string]interface{} , map[string]interface{}){
	jsonString := decode(base64JSON)

	inputData := extractInputData(jsonString)
	outputData := extractOutputData(jsonString)
	
	return inputData , outputData

}



// Function to extract and return input data as a map
func extractInputData(jsonString string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Extract the "input" section
	input, ok := data["input"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: 'input' is not a valid map structure")
		return nil
	}

	return input
}

// Function to extract and return output data as a map
func extractOutputData(jsonString string) map[string]interface{} {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Extract the "output" section
	output, ok := data["output"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: 'output' is not a valid map structure")
		return nil
	}

	return output
}

// // Function to get inner input parameter keys in order
// func innerInputParameter(jsonString string) []string {
// 	decoder := json.NewDecoder(strings.NewReader(jsonString))
// 	var (
// 		token   json.Token
// 		inInput bool
// 		keyOrder []string
// 		err     error
// 	)

// 	// Process the tokens
// 	for {
// 		token, err = decoder.Token()
// 		if err != nil {
// 			break
// 		}

// 		switch t := token.(type) {
// 		case string:
// 			if t == "input" {
// 				inInput = true
// 			} else if inInput {
// 				keyOrder = append(keyOrder, t)
// 			}
// 		case json.Delim:
// 			if t.String() == "}" && inInput {
// 				inInput = false
// 			}
// 		}
// 	}
// 	return keyOrder
// }

// // Function to get inner output parameter keys in order
// func innerOutputParameter(jsonString string) []string {
// 	decoder := json.NewDecoder(strings.NewReader(jsonString))
// 	var (
// 		token   json.Token
// 		inOutput bool
// 		keyOrder []string
// 		err     error
// 	)

// 	// Process the tokens
// 	for {
// 		token, err = decoder.Token()
// 		if err != nil {
// 			break
// 		}

// 		switch t := token.(type) {
// 		case string:
// 			if t == "output" {
// 				inOutput = true
// 			} else if inOutput {
// 				keyOrder = append(keyOrder, t)
// 			}
// 		case json.Delim:
// 			if t.String() == "}" && inOutput {
// 				inOutput = false
// 			}
// 		}
// 	}
// 	return keyOrder
// }




