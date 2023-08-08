package uploadbusiness

import (
	"context"
	"fmt"
	"go-api/common"
	"go-api/component/uploadpvd"
	uploadmodel "go-api/module/upload/model"
	"strings"
	"time"
)

type uploadFileBusiness struct {
	provider uploadpvd.UploadProvider
}

func NewUploadFileBusiness(provider uploadpvd.UploadProvider) *uploadFileBusiness {
	return &uploadFileBusiness{provider: provider}
}

func (business *uploadFileBusiness) UploadFile(context context.Context, data []byte, folder string) (*common.Image, error) {
	if strings.TrimSpace(folder) == "" {
		folder = "images"
	}

	fileName := fmt.Sprintf("%d", time.Now().Nanosecond())

	img, err := business.provider.SaveFileUploaded(context, data, folder, fileName)

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}
