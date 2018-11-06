package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"gopkg.in/loremipsum.v1"
)

var  loremGenerator = loremipsum.New()

func Handler(request events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{ Body: loremGenerator.Sentence(), StatusCode: 200 },nil
}

func main() {
	lambda.Start(Handler)
}
