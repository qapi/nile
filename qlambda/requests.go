package qlambda

import (
	"encoding/json"
	"errors"
)

// GetReqHeaders function returns map containing specific headers from lambda
// response and nil/error as second parameter
func GetReqHeaders(event map[string]interface{}, headers ...string) (map[string]interface{}, error) {

	resp := make(map[string]interface{})

	reqheaders, err := IfEventParamOK(event, "headers", "Request Headers")
	if err != nil {
		// TODO: check if this needs to return error
		return resp, nil
	}

	for _, name := range headers {
		value, err := IfReqParamEmptyOrMissing(reqheaders, name)
		if err != nil {
			return resp, err
		}
		resp[name] = value
	}

	return resp, nil

}

// GetReqBodyParam returns map with specific parameters from body of lambda
// response and nil/error as second parameter
func GetReqBodyParam(event map[string]interface{}, params ...string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})
	reqbody := make(map[string]interface{})

	bodystr, ok := event["body"].(string)
	if !ok {
		return resp, errors.New("invalid request, body is missing")
	}
	if err := json.Unmarshal([]byte(bodystr), &reqbody); err != nil {
		return resp, err
	}

	for _, name := range params {
		value, err := IfReqParamEmptyOrMissing(reqbody, name)
		if err != nil {
			return resp, err
		}
		resp[name] = value
	}

	return resp, nil

}

// GetReqIdentityParam function that gets identity parameter from lambda response
// returns map and nil/error
func GetReqIdentityParam(event map[string]interface{}, params ...string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})

	reqctx, err := IfEventParamOK(event, "requestContext", "Request Context")
	if err != nil {
		return resp, err
	}
	reqid, err := IfEventParamOK(reqctx, "identity", "Request Identity")
	if err != nil {
		return resp, err
	}

	for _, name := range params {
		value, err := IfReqParamEmptyOrMissing(reqid, name)
		if err != nil {
			return resp, err
		}
		resp[name] = value
	}

	return resp, nil

}

// GetQueryParam function for extracting query parameters from request made to
// lambda function, returns map and nil/error
func GetQueryParam(event map[string]interface{}, params ...string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})

	queryParams, err := IfEventParamOK(event, "queryStringParameters", "Query String Parameters")
	if err != nil {
		return resp, err
	}

	for _, name := range params {
		value, err := IfReqParamEmptyOrMissing(queryParams, name)
		if err != nil {
			return resp, err
		}
		resp[name] = value
	}

	return resp, nil
}
