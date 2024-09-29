package main

import (
	"bytes"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"back/codex"
)

// Define the ExecuteRequest struct to match the Judge0 API structure
type ExecuteRequest struct {
	LanguageID int    `json:"language_id"`
	SourceCode string `json:"source_code"`
	Stdin      string `json:"stdin,omitempty"`
	Base64     bool   `json:"base64_encoded"`
	Wait       bool   `json:"wait"`
}

// Define the structure to receive the token
type TokenResponse struct {
	Token string `json:"token"`
}

// Updated SubmissionResponse to match the Judge0 response format
type SubmissionResponse struct {
	Stdout         interface{} `json:"stdout"`
	Time           interface{} `json:"time"`
	Memory         interface{} `json:"memory"`
	Stderr         interface{} `json:"stderr"`
	Token          string      `json:"token"`
	CompileOutput  string      `json:"compile_output"`
	Message        interface{} `json:"message"`
	Status         Status      `json:"status"`
}

type Status struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {

	// ------------Golang--------------------------------
	// header:="cGFja2FnZSBtYWluCgppbXBvcnQgImZtdCIK"
	// class_code:=""
	// namespace:=""
	// func_code:="ZnVuYyB0d29TdW0obnVtcyBbXWludCwgdGFyZ2V0IGludCkgW11pbnQgewogICAgZm9yIGkgOj0gMDsgaSA8IGxlbihudW1zKTsgaSsrIHsKICAgICAgICBmb3IgaiA6PSBpICsgMTsgaiA8IGxlbihudW1zKTsgaisrIHsKICAgICAgICAgICAgaWYgbnVtc1tpXStudW1zW2pdID09IHRhcmdldCB7CiAgICAgICAgICAgICAgICByZXR1cm4gW11pbnR7aSwgan0KICAgICAgICAgICAgfQogICAgICAgIH0KICAgIH0KICAgIHJldHVybiBuaWwKfQ=="
	// main_code1:="ZnVuYyBtYWluKCkgew=="
	// test_case:="ewogICJpbnB1dCI6IHsKICAgICJudW1zIjogWwogICAgICBbMiwgNywgMTEsIDE1XSwKICAgICAgWzMsIDIsIDRdLAogICAgICBbMywgM10sCiAgICAgIFsxLCA1LCAzLCA2XSwKICAgICAgWzIsIDgsIDEyLCA1LCA3XSwKICAgICAgWzEwLCAxNSwgMiwgNywgOV0sCiAgICAgIFs1LCAzLCA4LCAxMiwgN10sCiAgICAgIFs0LCA5LCAxMSwgMTNdLAogICAgICBbNiwgMywgMiwgMSwgOF0sCiAgICAgIFswLCAwLCA0LCAwXQogICAgXSwKICAgICJ0YXJnZXQiOiBbOSwgNiwgNiwgNywgMTMsIDE3LCAxMCwgMjAsIDEwLCAwXQogIH0sCiAgIm91dHB1dCI6IHsKICAgICJvdXQiOiBbCiAgICAgIFswLCAxXSwgIAogICAgICBbMSwgMl0sICAKICAgICAgWzAsIDFdLCAgCiAgICAgIFswLCAzXSwgIAogICAgICBbMCwgNF0sICAKICAgICAgWzIsIDNdLCAgCiAgICAgIFswLCAxXSwgIAogICAgICBbMiwgM10sICAKICAgICAgWzAsIDRdLCAgCiAgICAgIFswLCAxXSAgCiAgICBdCiAgfQp9"
	// main_code2:="Zm9yIGksIG51bXMgOj0gcmFuZ2UgbnVtc19zIHsKICAgICAgICB0YXJnZXQgOj0gdGFyZ2V0X3NbaV0KICAgICAgICBleHBlY3RlZCA6PSBvdXRfc1tpXQogICAgICAgIHJlc3VsdCA6PSB0d29TdW0obnVtcywgdGFyZ2V0KQoKICAgICAgICBpZiAhY29tcGFyZVNsaWNlcyhyZXN1bHQsIGV4cGVjdGVkKSB7CiAgICAgICAgICAgIGZtdC5QcmludGYoIlRlc3QgY2FzZSAlZCBmYWlsZWQuXG4iLCBpKzEpCiAgICAgICAgICAgIGZtdC5QcmludGYoIklucHV0OiBudW1zID0gJXYsIHRhcmdldCA9ICVkXG4iLCBudW1zLCB0YXJnZXQpCiAgICAgICAgICAgIGZtdC5QcmludGYoIkV4cGVjdGVkOiAldiwgR290OiAldlxuXG4iLCBleHBlY3RlZCwgcmVzdWx0KQogICAgICAgIH0gZWxzZSB7CiAgICAgICAgICAgIGZtdC5QcmludGYoIlRlc3QgY2FzZSAlZDogQWNjZXB0ZWRcbiIsIGkrMSkKICAgICAgICB9CiAgICB9Cn0="
	// utility:="ZnVuYyBjb21wYXJlU2xpY2VzKGEsIGIgW11pbnQpIGJvb2wgewogICAgaWYgbGVuKGEpICE9IGxlbihiKSB7CiAgICAgICAgcmV0dXJuIGZhbHNlCiAgICB9CiAgICBmb3IgaSA6PSByYW5nZSBhIHsKICAgICAgICBpZiBhW2ldICE9IGJbaV0gewogICAgICAgICAgICByZXR1cm4gZmFsc2UKICAgICAgICB9CiAgICB9CiAgICByZXR1cm4gdHJ1ZQp9"
	// user_utility:=""


	// ----------------java------------------------------
	// header:="aW1wb3J0IGphdmEudXRpbC5BcnJheXM7CmltcG9ydCBqYXZhLnV0aWwuTGlzdDsKaW1wb3J0IGphdmEudXRpbC5zdHJlYW0uQ29sbGVjdG9yczs="
	// class_code:="cHVibGljIGNsYXNzIE1haW4gew=="
	// namespace:=""
	// func_code:="cHVibGljIHN0YXRpYyBpbnRbXSB0d29TdW0oaW50W10gbnVtcywgaW50IHRhcmdldCkgewogICAgICAgIGZvciAoaW50IGkgPSAwOyBpIDwgbnVtcy5sZW5ndGg7IGkrKykgewogICAgICAgICAgICBmb3IgKGludCBqID0gaSArIDE7IGogPCBudW1zLmxlbmd0aDsgaisrKSB7CiAgICAgICAgICAgICAgICBpZiAobnVtc1tpXSArIG51bXNbal0gPT0gdGFyZ2V0KSB7CiAgICAgICAgICAgICAgICAgICAgcmV0dXJuIG5ldyBpbnRbXXtpLCBqfTsKICAgICAgICAgICAgICAgIH0KICAgICAgICAgICAgfQogICAgICAgIH0KICAgICAgICByZXR1cm4gbnVsbDsKICAgIH0="
	// main_code1:="IHB1YmxpYyBzdGF0aWMgdm9pZCBtYWluKFN0cmluZ1tdIGFyZ3MpIHs="
	// test_case:="ewogICJpbnB1dCI6IHsKICAgICJudW1zIjogWwogICAgICBbMiwgNywgMTEsIDE1XSwKICAgICAgWzMsIDIsIDRdLAogICAgICBbMywgM10sCiAgICAgIFsxLCA1LCAzLCA2XSwKICAgICAgWzIsIDgsIDEyLCA1LCA3XSwKICAgICAgWzEwLCAxNSwgMiwgNywgOV0sCiAgICAgIFs1LCAzLCA4LCAxMiwgN10sCiAgICAgIFs0LCA5LCAxMSwgMTNdLAogICAgICBbNiwgMywgMiwgMSwgOF0sCiAgICAgIFswLCAwLCA0LCAwXQogICAgXSwKICAgICJ0YXJnZXQiOiBbOSwgNiwgNiwgNywgMTMsIDE3LCAxMCwgMjAsIDEwLCAwXQogIH0sCiAgIm91dHB1dCI6IHsKICAgICJvdXQiOiBbCiAgICAgIFswLCAxXSwgIAogICAgICBbMSwgMl0sICAKICAgICAgWzAsIDFdLCAgCiAgICAgIFswLCAzXSwgIAogICAgICBbMCwgNF0sICAKICAgICAgWzIsIDNdLCAgCiAgICAgIFswLCAxXSwgIAogICAgICBbMiwgM10sICAKICAgICAgWzAsIDRdLCAgCiAgICAgIFswLCAxXSAgCiAgICBdCiAgfQp9"
	// main_code2:="Zm9yIChpbnQgaSA9IDA7IGkgPCBudW1zX3Muc2l6ZSgpOyBpKyspIHsKICAgICAgICAgICAgaW50W10gbnVtcyA9IGxpc3RUb0FycmF5KG51bXNfcy5nZXQoaSkpOwogICAgICAgICAgICBpbnQgdGFyZ2V0ID0gdGFyZ2V0X3MuZ2V0KGkpOwogICAgICAgICAgICBpbnRbXSBleHBlY3RlZCA9IGxpc3RUb0FycmF5KG91dF9zLmdldChpKSk7CiAgICAgICAgICAgIGludFtdIHJlc3VsdCA9IHR3b1N1bShudW1zLCB0YXJnZXQpOwoKICAgICAgICAgICAgaWYgKCFjb21wYXJlU2xpY2VzKHJlc3VsdCwgZXhwZWN0ZWQpKSB7CiAgICAgICAgICAgICAgICBTeXN0ZW0ub3V0LnByaW50ZigiVGVzdCBjYXNlICVkIGZhaWxlZC5cbiIsIGkgKyAxKTsKICAgICAgICAgICAgICAgIFN5c3RlbS5vdXQucHJpbnRmKCJJbnB1dDogbnVtcyA9ICVzLCB0YXJnZXQgPSAlZFxuIiwgQXJyYXlzLnRvU3RyaW5nKG51bXMpLCB0YXJnZXQpOwogICAgICAgICAgICAgICAgU3lzdGVtLm91dC5wcmludGYoIkV4cGVjdGVkOiAlcywgR290OiAlc1xuXG4iLCBBcnJheXMudG9TdHJpbmcoZXhwZWN0ZWQpLCBBcnJheXMudG9TdHJpbmcocmVzdWx0KSk7CiAgICAgICAgICAgIH0gZWxzZSB7CiAgICAgICAgICAgICAgICBTeXN0ZW0ub3V0LnByaW50ZigiVGVzdCBjYXNlICVkOiBBY2NlcHRlZFxuIiwgaSArIDEpOwogICAgICAgICAgICB9CiAgICAgICAgfQogICAgfQp9"
	// utility:="ICBwdWJsaWMgc3RhdGljIGJvb2xlYW4gY29tcGFyZVNsaWNlcyhpbnRbXSBhLCBpbnRbXSBiKSB7CiAgICAgICAgaWYgKGEgPT0gbnVsbCB8fCBiID09IG51bGwgfHwgYS5sZW5ndGggIT0gYi5sZW5ndGgpIHsKICAgICAgICAgICAgcmV0dXJuIGZhbHNlOwogICAgICAgIH0KICAgICAgICBmb3IgKGludCBpID0gMDsgaSA8IGEubGVuZ3RoOyBpKyspIHsKICAgICAgICAgICAgaWYgKGFbaV0gIT0gYltpXSkgewogICAgICAgICAgICAgICAgcmV0dXJuIGZhbHNlOwogICAgICAgICAgICB9CiAgICAgICAgfQogICAgICAgIHJldHVybiB0cnVlOwogICAgfQ=="
	// user_utility:="ICBwdWJsaWMgc3RhdGljIGludFtdIGxpc3RUb0FycmF5KExpc3Q8SW50ZWdlcj4gbGlzdCkgewogICAgICAgIHJldHVybiBsaXN0LnN0cmVhbSgpLm1hcFRvSW50KGkgLT4gaSkudG9BcnJheSgpOwogICAgfQ=="

	testcase:=codex.DeclarationJava(test_case)


source:=accumulated(header,namespace,class_code,func_code,main_code1,main_code2,testcase,utility,user_utility)

	// Set up your request data
	requestPayload := ExecuteRequest{
		LanguageID: 62, // Go language ID in Judge0
		// SourceCode: "cGFja2FnZSBtYWluCgppbXBvcnQgKAoJImZtdCIKCSJyZWZsZWN0IgopCgovLyBUd29TdW0gZnVuY3Rpb24KZnVuYyBUd29TdW0obnVtcyBbXWludCwgdGFyZ2V0IGludCkgW11pbnQgewoJbSA6PSBtYWtlKG1hcFtpbnRdaW50KQoJZm9yIGksIG51bSA6PSByYW5nZSBudW1zIHsKCQlpZiB2YWwsIG9rIDo9IG1bdGFyZ2V0LW51bV07IG9rIHsKCQkJcmV0dXJuIFtdaW50e3ZhbCwgaX0KCQl9CgkJbVtudW1dID0gaQoJfQoJcmV0dXJuIG5pbAp9CgovLyBNYWluIGZ1bmN0aW9uIHRvIGhhbmRsZSB0ZXN0IGNhc2VzCmZ1bmMgbWFpbigpIHsKCS8vIFRlc3QgY2FzZXMKCXRlc3RDYXNlcyA6PSBbXXN0cnVjdCB7CgkJbnVtcyAgICAgICAgICAgW11bXWludCAvLyAyRCBhcnJheSBvZiBpbnB1dHMgZm9yIFR3b1N1bQoJCXRhcmdldHMgICAgICAgIFtdaW50ICAgLy8gMUQgYXJyYXkgb2YgdGFyZ2V0cwoJCWV4cGVjdGVkT3V0cHV0IFtdW11pbnQgLy8gMkQgYXJyYXkgb2YgZXhwZWN0ZWQgb3V0cHV0cwoJfXsKCQl7CgkJCW51bXM6IFtdW11pbnR7CgkJCQl7MiwgNywgMTEsIDE1fSwKCQkJCXszLCAyLCA0fSwKCQkJCXszLCAzfSwKCQkJfSwKCQkJdGFyZ2V0czogW11pbnR7OSwgNiwgNn0sCgkJCWV4cGVjdGVkT3V0cHV0OiBbXVtdaW50ewoJCQkJezAsIDF9LAoJCQkJezEsIDJ9LAoJCQkJezAsIDF9LAoJCQl9LAoJCX0sCgl9CgoJLy8gTG9vcCB0aHJvdWdoIHRlc3QgY2FzZXMKCWZvciBpLCB0ZXN0Q2FzZSA6PSByYW5nZSB0ZXN0Q2FzZXMgewoJCWZvciBqLCBudW1zIDo9IHJhbmdlIHRlc3RDYXNlLm51bXMgewoJCQl0YXJnZXQgOj0gdGVzdENhc2UudGFyZ2V0c1tqXQoJCQlleHBlY3RlZCA6PSB0ZXN0Q2FzZS5leHBlY3RlZE91dHB1dFtqXQoJCQlyZXN1bHQgOj0gVHdvU3VtKG51bXMsIHRhcmdldCkKCgkJCS8vIENoZWNrIGlmIHRoZSByZXN1bHQgbWF0Y2hlcyB0aGUgZXhwZWN0ZWQgb3V0cHV0CgkJCWlmICFyZWZsZWN0LkRlZXBFcXVhbChyZXN1bHQsIGV4cGVjdGVkKSB7CgkJCQlmbXQuUHJpbnRmKCJUZXN0IGNhc2UgJWQgZmFpbGVkOiBnb3QgJXYsIGV4cGVjdGVkICV2XG4iLCBpKzEsIHJlc3VsdCwgZXhwZWN0ZWQpCgkJCX0KCQl9Cgl9Cn0K",
		SourceCode:source,
		Base64:     true,
		Wait:       false,
	}

	// Convert requestPayload to JSON
	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Send the POST request
	resp, err := http.Post("http://192.168.1.12:2358/submissions/?base64_encoded=true&wait=false", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response to get the token
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		fmt.Println("Error unmarshalling token response:", err)
		return
	}

	fmt.Printf("Token received: %s\n", tokenResponse.Token)




	// Wait for 5 seconds before checking the status
	fmt.Println("Waiting for 2 seconds before checking execution status...")
	time.Sleep(2 * time.Second)






	// Check the submission status
	statusURL := fmt.Sprintf("http://192.168.1.12:2358/submissions/%s?base64_encoded=false", tokenResponse.Token)
	fmt.Printf("Checking status at URL: %s\n", statusURL)

	statusResp, err := http.Get(statusURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer statusResp.Body.Close()

	// Read the status response
	statusBody, err := ioutil.ReadAll(statusResp.Body)
	if err != nil {
		fmt.Println("Error reading status response:", err)
		return
	}

	


	var submissionRes interface{}

	// Unmarshal the JSON response
	err = json.Unmarshal(statusBody, &submissionRes)
	if err != nil {
		fmt.Println("Error unmarshalling submission response:", err)
		return
	}

	// Marshal it back to pretty JSON for printing
	prettyJSON, err := json.MarshalIndent(submissionRes, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling submission response:", err)
		return
	}

	// Print the formatted JSON
	fmt.Println(string(prettyJSON))


	var submissionResponse SubmissionResponse
	err = json.Unmarshal(statusBody, &submissionResponse)
	if err != nil {
		fmt.Println("Error unmarshalling submission response:", err)
		return
	}

	// Print the submission status and details
	fmt.Printf("Submission Status ID: %d \n", submissionResponse.Status.ID)
	fmt.Printf("Submission Status Description: %s \n", submissionResponse.Status.Description)


	fmt.Printf("Token: %s\n", submissionResponse.Token)
	
	if submissionResponse.Stdout != nil {
		fmt.Printf("Stdout: %v\n", submissionResponse.Stdout)
	}
	if submissionResponse.Stderr != nil {
		fmt.Printf("Stderr: %v\n", submissionResponse.Stderr)
	}
	if submissionResponse.CompileOutput != "" {
		fmt.Printf("Compile Output: %s\n", submissionResponse.CompileOutput)
	}
	if submissionResponse.Message != nil {
		fmt.Printf("Message: %v\n", submissionResponse.Message)
	}
	if submissionResponse.Time != nil {
		fmt.Printf("Execution Time: %v\n", submissionResponse.Time)
	}
	if submissionResponse.Memory != nil {
		fmt.Printf("Memory Used: %v\n", submissionResponse.Memory)
	}
}


func encode (s string) string{
	encodedString := base64.StdEncoding.EncodeToString([]byte(s))
	// fmt.Println("Encoded (Base64):", encodedString)
	return encodedString
}

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




func accumulated ( header string , namespace string,class_code string, func_code string , main_code1 string, main_code2 string , testcase string, utility string , user_utility string) string{
spaceTwo := "Cgo="
spaceOne := "Cg=="

	q :=decode(header)+decode(spaceTwo)+decode(namespace)+decode(spaceOne)+decode(class_code)+decode(spaceOne)+decode(func_code)+decode(spaceOne)+decode(utility)+decode(spaceOne)+decode(user_utility)+decode(spaceTwo)+decode(main_code1)+decode(spaceOne)+testcase+decode(spaceOne)+decode(main_code2)
	fmt.Println("accumulated function")
	fmt.Println(q)
	fmt.Println("accumulated function")
	
	z := encode(q)
	return z
}
	
	
