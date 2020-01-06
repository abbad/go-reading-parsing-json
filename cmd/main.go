package main

import (
	"fmt"
	"pkg/jsonweb"
)

func main() {
	url := "https://dog.ceo/api/breeds/list/all"

	response := jsonweb.ReadContentFromServer(url)

	decodedResponse := jsonweb.ResponseToJson(response)
	
	for dogType, breed := range decodedResponse.Message {
		// index is the index where we are
		// element is the element from someSlice for where we are
		fmt.Printf("Type: %v Breed: %v \n", dogType, breed)
	}
}
