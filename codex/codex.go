package codex

import (
	// "fmt"
	"strings"
	"back/init/golang"
	"back/init/handlerjson"
	"back/init/java"
	"back/init/javascript"
	"back/init/kotlin"
	"back/init/cpp"
	"back/init/dart"
	"back/init/dotnet"
	"back/init/python"
	"back/init/ruby"
	"back/init/swift"
	"back/init/rust"
	"back/init/typescript"
)


func DeclarationGo(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := golang.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := golang.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationJava(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := java.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := java.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}

func DeclarationJavascript(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := javascript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := javascript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}

func DeclarationKotlin(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := kotlin.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := kotlin.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}



func DeclarationCplusplus(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := cpp.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := cpp.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationDart(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := dart.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := dart.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationCSharp(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := dotnet.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := dotnet.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationPython(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := python.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := python.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationRuby(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := ruby.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := ruby.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationSwift(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := swift.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := swift.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}


func DeclarationRust(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := rust.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := rust.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}



func DeclarationTypescript(base64JSON string) (array1 string) {
	// Extract input and output maps from the JSON string
	inputMap, outputMap := handlerjson.ExtractDataMap(base64JSON)

	// Use strings.Builder for more efficient string concatenation
	var builder strings.Builder

	// Process the input map
	for key, value := range inputMap {
		array := typescript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n") // Add newline for each array
	}

	// Process the output map
	for key, value := range outputMap {
		array := typescript.ProcessJSON(value, key)
		builder.WriteString(array)
		builder.WriteString("\n")
	}

	// Convert the strings.Builder to a string
	array1 = builder.String()

	// Remove the trailing newline
	array1 = strings.TrimSuffix(array1, "\n")

	// fmt.Println(array1) // Optional: To print the final output
	return array1
}