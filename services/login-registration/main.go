package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/database"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/model/users"
	"github.com/radekkrejcirik01/Casblanca-backend/services/login-registration/pkg/rest"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	database.Connect()
	if err := database.DB.AutoMigrate(&users.User{}, &users.Photo{}, &users.Tag{}); err != nil {
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
