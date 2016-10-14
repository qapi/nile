package qlambda

import "net/http"

// KOResponse prepares response when error happens
// accepts error, returns body and status code that
// that should be passed to clent
func KOResponse(reason string) (body map[string]interface{}, statusCode int) {
	body = make(map[string]interface{})
	body["status"] = "ko"
	body["reason"] = reason
	statusCode = http.StatusBadRequest
	return body, statusCode

}
