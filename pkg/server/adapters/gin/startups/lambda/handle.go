package startup

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

func handle(request events.APIGatewayProxyRequest) (meta server.RequestMeta) {
	if request.RequestContext.Authorizer == nil {
		return
	}

	var session server.Session
	schema := server.TransformDataToSchema(session)
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           &schema,
		WeaklyTypedInput: true,
	})

	if err != nil {
		return
	}

	err = decoder.Decode(request.RequestContext.Authorizer)
	if err != nil {
		return
	}

	validate := validator.New()
	if err := validate.Struct(schema); err == nil {
		session = server.TransformSchemaToData[server.Session](schema)
		meta.Session = &session
	}

	return
}
