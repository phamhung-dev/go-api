package userbusiness

import (
	"context"
	"go-api/component/tokenpvd"
	usermodel "go-api/module/user/model"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type LoginStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	storage       LoginStorage
	tokenProvider tokenpvd.TokenProvider
	expiredIn     int
}

func NewLoginBusiness(storage LoginStorage, tokenProvider tokenpvd.TokenProvider) *loginBusiness {
	expiredIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_IN"))

	if err != nil {
		expiredIn = 60 * 60 * 24 * 30
	}

	return &loginBusiness{
		storage:       storage,
		tokenProvider: tokenProvider,
		expiredIn:     expiredIn,
	}
}

func (business *loginBusiness) Login(context context.Context, data *usermodel.UserLogin) (*tokenpvd.Token, error) {
	conditions := map[string]interface{}{
		"email": data.Email,
	}

	user, err := business.storage.Find(context, conditions)

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordIsIncorrect
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return nil, usermodel.ErrEmailOrPasswordIsIncorrect
	}

	payload := tokenpvd.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	token, err := business.tokenProvider.Generate(payload, business.expiredIn)

	if err != nil {
		return nil, err
	}

	return token, nil
}
