package startup

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
)

func handle(request events.APIGatewayProxyRequest) (meta server.RequestMeta) {
	if request.RequestContext.Authorizer == nil {
		return
	}

	session, err := common.DecodeStruct[server.Session](request.RequestContext.Authorizer)
	if err != nil {
		return
	}

	meta.Session = &session

	return
}
