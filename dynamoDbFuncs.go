package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
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
