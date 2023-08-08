package uploadpvd

import (
	"bytes"
	"context"
	"go-api/common"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	_ "github.com/joho/godotenv/autoload"
)

type cloudinaryProvider struct {
	apiKey    string
	apiSecret string
	cloudName string
}

func NewCloudinaryProvider() *cloudinaryProvider {
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")

	if apiKey == "" || apiSecret == "" || cloudName == "" {
		panic(ErrProviderIsNotConfigured)
	}

	return &cloudinaryProvider{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		cloudName: cloudName,
	}
}

func (provider *cloudinaryProvider) SaveFileUploaded(context context.Context, data []byte, folder string, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)

	cld, err := cloudinary.NewFromParams(provider.cloudName, provider.apiKey, provider.apiSecret)
	if err != nil {
		return nil, err
	}

	if uploadResult, err := cld.Upload.Upload(context, fileBytes, uploader.UploadParams{
		PublicID: fileName,
		Folder:   folder,
	}); err != nil {
		return nil, err
	} else {
		return &common.Image{
			ID:        uploadResult.PublicID,
			Url:       uploadResult.SecureURL,
			Width:     uploadResult.Width,
			Height:    uploadResult.Height,
			CloudName: "cloudinary",
			Extension: uploadResult.Format,
		}, nil
	}
}
