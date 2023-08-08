package userbusiness

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
)

type UpdateUserStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Update(context context.Context, id uuid.UUID, data *usermodel.UserUpdate) (*usermodel.User, error)
}

type updateUserBusiness struct {
	storage UpdateUserStorage
}

func NewUpdateUserBusiness(storage UpdateUserStorage) *updateUserBusiness {
	return &updateUserBusiness{storage: storage}
}

func (business *updateUserBusiness) UpdateUser(context context.Context, id uuid.UUID, data *usermodel.UserUpdate) (*usermodel.User, error) {
	conditions := map[string]interface{}{
		"id": id,
	}

	user, err := business.storage.Find(context, conditions)

	if err == common.ErrRecordNotFound {
		return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if err != nil {
		return nil, common.ErrDB(err)
	}

	if user.IsLocked {
		return nil, usermodel.ErrUserIsLocked
	}

	if err := data.Validate(); err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	user, err = business.storage.Update(context, id, data)

	if err != nil {
		return nil, common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	user.Mask(false)

	return user, nil
}
