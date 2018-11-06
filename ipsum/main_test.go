package main

import (
	"testing"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {


	tests := []struct {
		request events.APIGatewayProxyRequest
		expect string
		err error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: "test"},
			expect: "Go Serverless v1.0! Your function executed successfully!",
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err,err)
		assert.NotEmpty(t,response.Body)
		assert.Equal(t,response.StatusCode,200)
	}
}
