package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

func getAllDynamoDBDocs() ([]CmpResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	param := &dynamodb.ScanInput{
		TableName: aws.String("nwolf-top10-cmp"),
	}
	scan, err := client.Scan(context.TODO(), param)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	var obj []CmpResponse
	var dec attributevalue.DecoderOptions
	dec.TagKey = "json"
	err = attributevalue.UnmarshalListOfMapsWithOptions(scan.Items, &obj)
	if err != nil {
		log.Fatalf("unable to unmarshal records: %v", err)
	}
	return obj, err
}

func getDocumentCount() (int32, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return 0, err
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	param := &dynamodb.ScanInput{
		TableName: aws.String("nwolf-top10-cmp"),
	}
	scan, err := client.Scan(context.TODO(), param)
	if err != nil {
		return 0, err
	}
	return scan.Count, nil
}

func getDocDateRange(lower string, upper string) ([]CmpResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)
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
		log.Fatalf("Query API call failed: %s", err)
	}
	var obj []CmpResponse
	err = attributevalue.UnmarshalListOfMapsWithOptions(response.Items, &obj)
	if err != nil {
		log.Fatalf("unable to unmarshal records: %v", err)
	}
	return obj, err
}
