package ginupload

import (
	"go-api/common"
	"go-api/component/appctx"
	uploadbusiness "go-api/module/upload/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(appctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "images")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uploadProvider := appctx.GetUploadProvider()
		business := uploadbusiness.NewUploadFileBusiness(uploadProvider)
		img, err := business.UploadFile(c.Request.Context(), dataBytes, folder)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
