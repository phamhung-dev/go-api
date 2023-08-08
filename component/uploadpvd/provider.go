package uploadpvd

import (
	"context"
	"errors"
	"go-api/common"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, folder string, fileName string) (*common.Image, error)
}

var (
	ErrProviderIsNotConfigured = common.NewCustomErrorResponse(
		errors.New("upload provider is not configured"),
		"upload provider is not configured",
		"ErrProviderIsNotConfigured",
	)
)
