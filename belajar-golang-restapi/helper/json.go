package helper

import (
	"net/http"
	"encoding/json"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	
	// Set HTTP status code based on response code if it's a WebResponse
	if webResp, ok := response.(interface{ GetCode() int }); ok {
		writer.WriteHeader(webResp.GetCode())
	}
	
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
