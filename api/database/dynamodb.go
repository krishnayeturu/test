package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dbClient struct{ *dynamodb.Client }

func CreateDynamoDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "AK_ID", SecretAccessKey: "SECRET_KEY", SessionToken: "TOKEN",
			},
		}),
		// config.WithEndpointResolverWithOptions(aws.EndpointResolverFunc(
		// 	func(service, region string) (aws.Endpoint, error) {
		// 		return aws.Endpoint{URL: "http://localhost:8000"}, nil
		// 	})),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}
