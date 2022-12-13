package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/radekkrejcirik01/Casblanca-backend/services/pushnotifications/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/pushnotifications/pkg/rest"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	database.Connect()

	fiberLambda = fiberadapter.New(rest.Create())
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, request)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
