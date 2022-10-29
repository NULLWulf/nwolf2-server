package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getAllDynamoDBDocs() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))
	svc := dynamodb.New(sess)

	resp, _ := svc.Scan(
		&dynamodb.ScanInput{
			TableName: aws.String("nwolf-top10-cmp"),
		})
	var obj []CmpResponse
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &obj)
	fmt.Printf("%v\n", obj)
}

//func getDocumentCount() (int32, error) {
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
//		return 0, err
//	}
//
//	return out.Count, nil
//}
