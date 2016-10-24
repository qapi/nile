package qlambda

import (
	"encoding/json"
	"errors"
)

func GetReqHeaders(event map[string]interface{}, headers ...string) (map[string]interface{}, error) {

	resp := make(map[string]interface{})

	reqheaders, err := IfEventParamOK(event, "headers", "Request Headers")
	if err != nil {
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

func GetReqBodyParam(event map[string]interface{}, params ...string) (map[string]interface{}, error) {
	resp := make(map[string]interface{})
	reqbody := make(map[string]interface{})

	bodystr, ok := event["body"].(string)
	if !ok {
		return resp, errors.New("Invalid request! Body is missing.")
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
