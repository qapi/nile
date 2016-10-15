package qlambda

import (
	"errors"
	"net/http"

	"github.com/vsrc/nile/qlambda"
)

// KOResponse prepares response when error happens accepts error, returns body
// and status code that that should be passed to clent
func KOResponse(reason string) (body map[string]interface{}, statusCode int) {
	body = make(map[string]interface{})
	body["status"] = "ko"
	body["reason"] = reason
	statusCode = http.StatusBadRequest
	return body, statusCode

}

// IfReqParamEmptyOrMissing helper for lambda function that is accessed through API Gateway and
// have query parameters accepts queryparameters map of string interface and name of
// parameter that we are looking for and returns string if parameter is found and not empty
// and nil error as a second return parameter from this function
func IfReqParamEmptyOrMissing(queryParams map[string]interface{}, paramName string) (string, error) {

	// check if parameter exist
	if _, ok := queryParams[paramName]; !ok {
		return "", errors.New("Missing " + paramName + " parameter which is required")
	}

	// try conversion to string
	value, ok := queryParams[paramName].(string)

	if !ok || value == "" {
		return "", errors.New("Empty " + paramName + " parameter which is required")
	}

	return value, nil
}

// IfEventParamOK helper for lambda function accepts event that is passed by lambda
// as a map of string interface, name of parameter and label of parameter that we are
// looking for and returns map of string interface if parameter is found, not empty,
// not malformed and nil error as a second return parameter from this function
func IfEventParamOK(event map[string]interface{}, paramName string, paramLabel string) (map[string]interface{}, error) {

	// check if missing parameter
	if event[paramName] == nil {
		return map[string]interface{}{}, errors.New("Missing " + paramLabel)
	}

	// check if parameter empty or malformed
	param, paramOK := event[paramName].(map[string]interface{})
	if !paramOK {
		return qlambda.KOResponse("Malformed " + paramLabel)
	}

	return param, nil
}
