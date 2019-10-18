package proxy

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NewResponse returns a valid APIGatewayProxyResponse object.
func NewResponse(
	err error,
	body string,
	cors ...bool,
) events.APIGatewayProxyResponse {
	headers := map[string]string{}

	if len(cors) > 0 && cors[0] {
		headers["Access-Control-Allow-Origin"] = "*"
	}

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       body,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       body,
	}
}
