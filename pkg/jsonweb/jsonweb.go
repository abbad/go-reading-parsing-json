package jsonweb

import (
	"io/ioutil"
	"net/http"	
	"encoding/json"
	"strings"	
)
	
type ApiResponse struct {
	Status string
	Message map[string][]string
}

func ReadContentFromServer(jsonURL string) string {
	resp, err := http.Get(jsonURL)
	checkError(err)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func ResponseToJson(content string) ApiResponse {
	apiResponse := ApiResponse{}

	decodor := json.NewDecoder(strings.NewReader(content))

	for decodor.More() {
		err := decodor.Decode(&apiResponse)
		checkError(err)
	}

	return apiResponse
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
