package error

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/graphql-go/graphql/gqlerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ferza17/event-driven-api-gateway/utils"
)

type Error struct {
	StatusCode int
	Error      error
}

func HandlerGrpcError(rawError error) error {
	var (
		mapError = map[codes.Code]error{
			// Internal section
			codes.Unknown:  xerrs.Mask(rawError, utils.ErrInternalServerError),
			codes.Internal: xerrs.Mask(rawError, utils.ErrInternalServerError),
			// not found section
			codes.NotFound: xerrs.Mask(rawError, utils.ErrNotFound),
			// bad request section
			codes.InvalidArgument: xerrs.Mask(rawError, utils.ErrBadRequest),
		}
	)
	grpcStatusError, _ := status.FromError(rawError)
	err, ok := mapError[grpcStatusError.Code()]
	if !ok {
		err = xerrs.Mask(err, utils.ErrInternalServerError)
	}
	return err
}

func HandleGraphQLError(graphQLError []gqlerrors.FormattedError) Error {
	var (
		mapError = map[string]Error{
			// Internal section
			utils.ErrInternalServerError.Error(): {
				StatusCode: http.StatusInternalServerError,
				Error:      utils.ErrInternalServerError,
			},
			// not found section
			utils.ErrNotFound.Error(): {
				StatusCode: http.StatusNotFound,
				Error:      utils.ErrNotFound,
			},
			// bad request section
			utils.ErrBadRequest.Error(): {
				StatusCode: http.StatusBadRequest,
				Error:      utils.ErrBadRequest,
			},
			utils.ErrItemOutOfStock.Error(): {
				StatusCode: http.StatusBadRequest,
				Error:      utils.ErrItemOutOfStock,
			},
			utils.ErrItemAlreadyExist.Error(): {
				StatusCode: http.StatusBadRequest,
				Error:      utils.ErrItemAlreadyExist,
			},
		}
		err Error
		ok  bool
	)
	for _, formattedError := range graphQLError {
		err, ok = mapError[formattedError.Message]
		if !ok {
		}
		return err
	}
	return err
}
