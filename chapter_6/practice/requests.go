package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://pokeapi.co/api/v2/pokemon/charizard"

func main() {
	fmt.Println("Network Requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// When you do res.Status, Go automatically dereferences the res pointer to access the Status field of the struct that res points to.
	// This is equivalent to (*res).Status, but Go allows you to omit the explicit dereference.
	fmt.Printf("Response Status: %v\n", res.Status)
	defer res.Body.Close() // closes network resources

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	// since pokemon is a pointer when we print it we need to dereference it with *
	pokemon, err := pokemonFromJson(content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *pokemon)

	fmt.Println("\nChallenge:")
	challenge()
}

// Example function to unmarshal JSON into the struct
func pokemonFromJson(content string) (*Pokemon, error) {
	pokemonJSON := new(pokemonJSON) // new makes a pointer to the struct

	// Unmarshal requires a byte slice, which means the data must be fully available in memory.
	// NewDecoder works with any io.Reader, allowing for more flexible and potentially more efficient input handling, especially with streams or large files.
	// In summary, the choice between json.NewDecoder and json.Unmarshal depends on the source of your JSON data and your application's memory efficiency requirements.

	decoder := json.NewDecoder(strings.NewReader(content))
	err := decoder.Decode(pokemonJSON) // Decode the JSON directly into the struct.
	// err := json.Unmarshal([]byte(content), &pokemonJSON)
	if err != nil {
		// return the memory address of an empty Pokemon struct and the error
		return &Pokemon{}, err
	}

	// Simplify the Types field to just a slice of strings for Pokemon types
	simplifiedTypes := make([]string, len(pokemonJSON.Types))
	for i, t := range pokemonJSON.Types {
		simplifiedTypes[i] = t.Type.Name
	}
	// return the memory address of a Pokemon struct with the Name and Type fields set and the error
	return &Pokemon{pokemonJSON.Name, simplifiedTypes}, nil
}

// pokemonJSON is a struct for defining parsing of the JSON response
type pokemonJSON struct {
	Name  string     `json:"name"` // Maps the JSON field "name" to the struct field Name.
	Types []struct { // A slice of anonymous structs.
		Type struct { // An anonymous struct inside the slice.
			Name string `json:"name"` // Maps the JSON field "name" inside "type" to the struct field Name.
		} `json:"type"` // Maps the JSON field "type" to the struct field Type.
	} `json:"types"` // Maps the JSON field "types" to the struct field Types.
}

// Pokemon is a struct to hold the data we care about
type Pokemon struct {
	Name string
	Type []string
}

// the function below returns the actual struct not a pointer to the struct (written before I tested out other ways)

// In Go, you can perform a type check (or type assertion) by using a period followed by the type you want to assert to, enclosed in parentheses.
// The syntax looks like this:
//     `value, ok := unknownTypeVariable.(TypeYouWantToAssert)`
// `unknownTypeVariable` is the interface value you're working with.
// `TypeYouWantToAssert` is the concrete type you believe `unknownTypeVariable` may hold.
// value will be the value of `unknownTypeVariable` asserted to `TypeYouWantToAssert` if the assertion is successful.
// ok is a boolean that will be true if the assertion succeeded and false otherwise.
// This mechanism allows you to safely check if an interface value is of a specific type before working with it as that type,
// thus avoiding runtime panics that would occur from incorrect type assertions.

// func pokemonFromJson(content string) Pokemon {
//     var result interface{}

// 	err := json.Unmarshal([]byte(content), &result) // []byte(content) converts the string to a byte slice
//     if err != nil {
//         panic(err) // Handle error appropriately
//     }

// 	var pokemon Pokemon

// 	// type-assert the result to a map[string]interface{}
//     if dataMap, ok := result.(map[string]interface{}); ok {

// 		// now check if "name" exists in the map and if it does, type-assert it to a string
//         if name, found := dataMap["name"].(string); found {
// 			pokemon.Name = name
// 		}

// 		// now check if "types" exists in the map and if it does:
// 		// type-assert it to an array of interfaces
// 		if typesArray, found := dataMap["types"].([]interface{}); found {
// 			// iterate over the array of interfaces
// 			for _, typesArrayEntry := range typesArray {
// 				// 3 checks chained:
// 				// Each entry in the array `types` is a dict[str, dict]
// 				// Each entry has a key = "type" with a value of dict[str, dict]
// 				// The dict stored at "type" has a key = "name" with a value of string
// 				if typeName, found := typesArrayEntry.(map[string]interface{})["type"].(map[string]interface{})["name"].(string); found {
// 					pokemon.Type = append(pokemon.Type, typeName)
// 				}
// 			}
// 		}

// 	}

// 	return pokemon
// }
