package utils

import (
	"encoding/json"
	"net/http"
)

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

var logger Logger

func ReadRequest(req *http.Request, obj interface{}) (err error) {
	logger.Log("Populating object")
	err = json.NewDecoder(req.Body).Decode(obj)
	if err != nil {
		logger.LogErrorObj("Failed to decode request", obj)
	}

	return nil
}

func WriteResponse(writer http.ResponseWriter, obj interface{}) (err error) {
	logger.Log("Populating object")

	err = json.NewEncoder(writer).Encode(obj)
	if err != nil {
		logger.LogErrorObj("Failed to write in response", err)
		return err
	}

	return nil
}
