package error

import (
	"net/http"

	"github.com/graphql-go/graphql/gqlerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type ErrorRest struct {
	Code int
	Err  error
}

func HandlerGraphQLErrorFromGRPC(errors []gqlerrors.FormattedError) (errRest ErrorRest) {
	var (
		mapError = map[codes.Code]ErrorRest{
			codes.Internal: {
				Code: http.StatusInternalServerError,
				Err:  utils.ErrInternalServerError,
			},
			codes.NotFound: {
				Code: http.StatusNotFound,
				Err:  utils.ErrNotFound,
			},
		}
		ok bool
	)

	for _, formattedError := range errors {
		grpcStatusError, _ := status.FromError(formattedError)
		if errRest, ok = mapError[grpcStatusError.Code()]; !ok {
			errRest = ErrorRest{
				Code: http.StatusInternalServerError,
				Err:  formattedError.OriginalError(),
			}
		}
		break
	}
	return
}
