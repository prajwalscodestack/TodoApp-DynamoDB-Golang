package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DBClient is a global DynamoDB client
var DBClient *dynamodb.Client

// InitDynamoDB initializes the DynamoDB client
func InitDynamoDB() {
	// Load AWS configuration (uses ~/.aws/credentials or IAM Role if running in AWS)
	cfg, err := config.LoadDefaultConfig(context.TODO()) //config.WithRegion("ap-south-1"), // Change region if needed

	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	// Create DynamoDB client
	DBClient = dynamodb.NewFromConfig(cfg)

	log.Println("DynamoDB client initialized successfully!")
}
