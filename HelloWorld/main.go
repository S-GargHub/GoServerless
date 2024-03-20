package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Input struct {
	Name string `json:"name"`
}

func handler(ctx context.Context, name Input) (string, error) {
	return fmt.Sprintf("Hello, %s!", name.Name), nil
}

func main() {
	lambda.Start(handler)
}