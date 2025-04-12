package meta

import (
	"context"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"google.golang.org/grpc/metadata"
)

func getMetadata(ctx context.Context) metadata.MD {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.Pairs()
	}

	return md
}

func SetOptionsContext(ctx context.Context, options Options) context.Context {
	ctx = context.WithValue(ctx, keyTimeZone, options.TimeZone)

	md := getMetadata(ctx)
	md.Set(string(keyTimeZone), options.TimeZone)

	return metadata.NewOutgoingContext(ctx, md)
}

func SetOptionsMetadataContext(ctx context.Context, md metadata.MD) (context.Context, error) {
	data, err := common.DecodeStruct[optionsMetadata](md)
	if err != nil {
		return nil, err
	}

	options, err := common.DecodeStruct[Options](map[string]any{
		string(keyTimeZone): data.TimeZone[0],
	})

	if err != nil {
		return nil, err
	}

	return SetOptionsContext(ctx, options), nil
}

func SetSessionHeader(session Session) map[string]any {
	data, _ := common.EncodeStruct(session)

	return data
}

func SetSessionContext(ctx context.Context, session Session) context.Context {
	ctx = context.WithValue(ctx, keySession, session)

	md := getMetadata(ctx)
	md.Set(string(keyUserID), session.UserID.String())

	return metadata.NewOutgoingContext(ctx, md)
}

func SetSessionMetadataContext(ctx context.Context, md metadata.MD) (context.Context, error) {
	data, err := common.DecodeStruct[sessionMetadata](md)
	if err != nil {
		return nil, err
	}

	session, err := common.DecodeStruct[Session](map[string]any{
		string(keyUserID): data.UserID[0],
	})

	if err != nil {
		return nil, err
	}

	return SetSessionContext(ctx, session), nil
}
