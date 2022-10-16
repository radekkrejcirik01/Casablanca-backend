package main

import (
	"context"
	"log"
	"login-registration/pkg/database"
	"login-registration/pkg/model/users"
	"login-registration/pkg/rest"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	database.Connect()
	if err := database.DB.AutoMigrate(&users.USERS{}); err != nil {
		log.Fatal(err)
	}
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
