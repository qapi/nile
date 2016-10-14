package qlambda

import "net/http"

// ErrorResponse prepares response when error happens
// accepts error, returns body and status code that
// that should be passed to clent
func ErrorResponse(err error) (body map[string]interface{}, statusCode int) {
	body = make(map[string]interface{})
	body["status"] = "ko"
	body["reason"] = err.Error
	statusCode = http.StatusBadRequest
	return body, statusCode

}
