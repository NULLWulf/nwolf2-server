package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

//func getAllDynamoDBDocs() {
//	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
//		o.Region = "us-east-1"
//		return nil
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	svc := dynamodb.NewFromConfig(cfg)
//
//	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
//		TableName: aws.String("nwolf-top10-cmp"),
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	var results []CmpResponse
//	var error = dynamodbattribute.UnmarshalListOfMaps(out.Count(), &results)
//
//	fmt.Println()
//}

func getDocumentCount() (int32, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("nwolf-top10-cmp"),
	})
	if err != nil {
		return 0, err
	}

	return out.Count, nil
}
