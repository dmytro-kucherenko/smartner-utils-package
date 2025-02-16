package startup

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func Init(create func(meta server.RequestMeta) *server.StartupOptions[gin.Engine]) {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		meta := handle(request)
		options := create(meta)

		return ginadapter.New(options.Router).ProxyWithContext(ctx, request)
	})
}
