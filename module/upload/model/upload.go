package uploadmodel

import "go-api/common"

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomErrorResponse(err, "cannot save file", "ErrCannotSaveFile")
}
