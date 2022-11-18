package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getAllDynamoDBDocs() ([]CmpResponse, error) {
	client, err := configDynamo()
	if err != nil {
		return nil, err
	}
	param := &dynamodb.ScanInput{
		TableName: aws.String("nwolf-top10-cmp"),
	}
	scan, err := client.Scan(context.TODO(), param)
	if err != nil {
		logInternalError(InternalError{
			Context: "getAllDynamoDBDocs() - scan",
			Error:   err.Error(),
		})
		return nil, err
	}
	var obj []CmpResponse
	err = attributevalue.UnmarshalListOfMapsWithOptions(scan.Items, &obj)
	if err != nil {
		logInternalError(InternalError{
			Context: "getAllDynamoDBDocs() - UnmarshallList",
			Error:   err.Error(),
		})
		return nil, err
	}
	return obj, err
}

func getDocumentCount() (int32, error) {
	client, err := configDynamo()
	if err != nil {
		return 0, err
	}
	param := &dynamodb.ScanInput{
		TableName: aws.String("nwolf-top10-cmp"),
	}
	scan, err := client.Scan(context.TODO(), param)
	if err != nil {
		logInternalError(InternalError{
			Context: "getDocumentCount() - scan error",
			Error:   err.Error(),
		})
		return 0, err
	}
	return scan.Count, err
}

func getDocDateRange(lower string, upper string) ([]CmpResponse, error) {
	client, err := configDynamo()
	if err != nil {
		return nil, err
	}
	var response *dynamodb.QueryOutput
	//keyEx := expression.KeyBetween(expression.Key("TimeBlockUTC"), expression.Value(lower), expression.Value(upper))
	keyEx := expression.KeyAnd(
		expression.Key("Partition").Equal(expression.Value("Top10Cryptos")),
		expression.Key("TimeBlockUTC").Between(expression.Value(lower), expression.Value(upper)))

	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	response, err = client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:                 aws.String("nwolf-top10-cmp"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		logInternalError(InternalError{
			Context: "getDocDate() - Error calling dynamodb.Query API",
			Error:   err.Error(),
		})
	}
	var obj []CmpResponse
	err = attributevalue.UnmarshalListOfMapsWithOptions(response.Items, &obj)
	if err != nil {
		logInternalError(InternalError{
			Context: "getDocRangeDate() - UnmarshallList",
			Error:   err.Error(),
		})
	}
	return obj, err
}

func configDynamo() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		logInternalError(InternalError{
			Context: "DynamoDB Configuration",
			Error:   err.Error(),
		})
		return nil, err
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)
	return client, err
}
