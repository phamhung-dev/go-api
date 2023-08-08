package userbusiness

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"
)

type CreateUserStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(context context.Context, data *usermodel.UserCreate) (*usermodel.User, error)
}

type createUserBusiness struct {
	storage CreateUserStorage
}

func NewCreateUserBusiness(storage CreateUserStorage) *createUserBusiness {
	return &createUserBusiness{storage: storage}
}

func (business *createUserBusiness) CreateUser(context context.Context, data *usermodel.UserCreate) (*usermodel.User, error) {
	conditions := map[string]interface{}{
		"email": data.Email,
	}

	_, err := business.storage.Find(context, conditions)

	if err == nil {
		// if user.IsLocked {
		// 	return nil, common.ErrEntityIsLocked(usermodel.EntityName, nil)
		// }

		return nil, usermodel.ErrEmailExisted
	}

	if err != common.ErrRecordNotFound {
		return nil, common.ErrDB(err)
	}

	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	user, err := business.storage.Create(context, data)

	if err != nil {
		return nil, common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	user.Mask(false)

	return user, nil
}
