package qlambda

import (
	"errors"
	"net/http"
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

// IfParamEmptyOrMissing helper for lambda function that is accessed through API Gateway and
// have query parameters accepts queryparameters map of string interface and name of
// parameter that we are looking for and returns string if parameter is found and not empty
// and nil error as a second return parameter from this function
func IfParamEmptyOrMissing(queryParams map[string]interface{}, paramName string) (string, error) {

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
