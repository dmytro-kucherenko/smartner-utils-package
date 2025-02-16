package startup

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/go-playground/validator/v10"
)

func handle(request events.APIGatewayProxyRequest) (meta server.RequestMeta) {
	validate := validator.New()

	if request.RequestContext.Authorizer == nil {
		return
	}

	var session server.Session
	schema := server.TransformDataToSchema(session)

	data, err := json.Marshal(request.RequestContext.Authorizer)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &schema)
	if err != nil {
		return
	}

	if err := validate.Struct(schema); err == nil {
		session = server.TransformSchemaToData[server.Session](schema)
		meta.Session = &session
	}

	return
}
