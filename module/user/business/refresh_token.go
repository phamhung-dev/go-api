package userbusiness

import (
	"context"
	"go-api/common"
	"go-api/component/tokenpvd"
	usermodel "go-api/module/user/model"
	"os"
	"strconv"
)

type RefreshTokenStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type refreshTokenBusiness struct {
	storage       RefreshTokenStorage
	tokenProvider tokenpvd.TokenProvider
	expiredIn     int
}

func NewRefreshTokenBusiness(storage RefreshTokenStorage, tokenProvider tokenpvd.TokenProvider) *refreshTokenBusiness {
	expiredIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_IN"))

	if err != nil {
		expiredIn = 60 * 60 * 24 * 30
	}

	return &refreshTokenBusiness{
		storage:       storage,
		tokenProvider: tokenProvider,
		expiredIn:     expiredIn,
	}
}

func (business *refreshTokenBusiness) RefreshToken(context context.Context, refreshToken string) (*tokenpvd.Token, error) {
	payload, err := business.tokenProvider.ValidateRefreshToken(refreshToken)

	if err != nil {
		return nil, err
	}

	conditions := map[string]interface{}{
		"id": payload.UserId,
	}

	user, err := business.storage.Find(context, conditions)

	if err == common.ErrRecordNotFound {
		panic(common.ErrEntityNotFound(usermodel.EntityName, err))
	}

	if err != nil {
		panic(common.ErrDB(err))
	}

	newPayload := tokenpvd.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	token, err := business.tokenProvider.Generate(newPayload, business.expiredIn)

	if err != nil {
		return nil, err
	}

	return token, nil
}
