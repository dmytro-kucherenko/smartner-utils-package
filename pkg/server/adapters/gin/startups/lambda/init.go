package startup

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/dtos"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

func Init(create func(logger types.Logger, meta server.RequestMeta) (adapter.StartupOptions, error)) {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		logger := log.New("Init")
		meta := handle(request)
		options, err := create(logger, meta)

		if err != nil {
			var response dtos.ErrorResponse
			httpErr, ok := err.(*errors.HttpError)

			if ok {
				response = dtos.ErrorResponse{
					Status:  httpErr.Status(),
					Message: httpErr.Error(),
					Details: httpErr.Details(),
				}
			} else {
				response = dtos.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "internal server error",
				}

				logger.Fatal(err.Error())
			}

			body, _ := json.Marshal(response)

			return events.APIGatewayProxyResponse{
				StatusCode: response.Status,
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       string(body),
			}, nil
		}

		return ginadapter.New(options.Router).ProxyWithContext(ctx, request)
	})
}
