package dyndb

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getSession(region string) *dynamodb.DynamoDB {
	return dynamodb.New(session.New(&aws.Config{Region: aws.String(region)}))
}

// QueryDB function abstracts query process of AWS DynamoDB
func QueryDB(region string, tableName string, appID string, keyName string, keyValue string) (*dynamodb.QueryOutput, error) {
	svc := getSession(region)

	var queryInput = &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("AppID = :appID AND #" + keyName + " = :" + keyName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":appID": &dynamodb.AttributeValue{
				S: aws.String(appID),
			},
			":" + keyName: &dynamodb.AttributeValue{
				S: aws.String(keyValue),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#" + keyName: aws.String(keyName),
		},
	}

	dbresp, err := svc.Query(queryInput)
	if err != nil {
		return nil, err
	}

	return dbresp, nil
}

// FlattenDBResponse function recursively loops through response from dynamodb output
// and flattens it to map with string as key and interface as value, for each parameter in output
// function is going through all provided types and discards all except one that is not nil
// and it formats it to matching type: string, integer, float or boolean
// if parameter does not match to one of these types all are returned
func FlattenDBResponse(resp *dynamodb.QueryOutput) (map[string]interface{}, error) {
	var v = make(map[string]interface{})
	// loop through series of versions of results
	for _, item := range resp.Items {
		// loop through array of parameters
		for paramKey, paramVal := range item {
			// flatten parameter value
			val, err := flattenParam(paramVal)
			if err != nil {
				return nil, err
			}
			v[paramKey] = val
		}
	}

	return v, nil
}

func flattenParam(paramVal *dynamodb.AttributeValue) (interface{}, error) {
	var v interface{}

	if paramVal.S != nil {
		// type is string
		v = paramVal.S
	} else if paramVal.N != nil {
		// type is number
		var num interface{}
		var err error

		// check if is float or int
		if strings.Contains(*paramVal.N, ".") {
			// it's a float
			num, err = strconv.ParseFloat(*paramVal.N, 10)
		} else {
			// it's an integer
			num, err = strconv.Atoi(*paramVal.N)
		}

		if err != nil {
			return nil, err
		}

		v = num

	} else if paramVal.BOOL != nil {
		// manage boolean type
		v = paramVal.BOOL
	} else if paramVal.M != nil {
		// manage nested maps
		var subMap = make(map[string]interface{})
		for subParamKey, subParamVal := range paramVal.M {
			subVal, err := flattenParam(subParamVal)
			if err != nil {
				return nil, err
			}

			subMap[subParamKey] = subVal
		}
		v = subMap
	} else {
		// no match, gracefully show all options
		v = paramVal
	}

	return v, nil
}
