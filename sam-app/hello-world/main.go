package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type dependency struct {
	ddb   dynamodbiface.DynamoDBAPI
	table string
}

// Record represents one record in the DynamoDB table
type Record struct {
	ID   string
	Body string
}

func (d *dependency) LambdaHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Check for dependencies, e.g., test injections.
	// If not present, create them with a live session.
	if d.ddb == nil {
		sess := session.Must(session.NewSession())
		svc := dynamodb.New(sess)

		d = &dependency{
			ddb:   svc,
			table: os.Getenv("DYNAMODB_TABLE"),
		}
	}

	// Create a new record from the request
	r := Record{
		ID:   request.RequestContext.RequestID,
		Body: request.Body,
	}

	// Marshal that record into a DynamoDB AttributeMap
	av, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Save the AttributeMap in the given table
	_, err = d.ddb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(d.table),
		Item:      av,
	})

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	d := dependency{}

	lambda.Start(d.LambdaHandler)
}