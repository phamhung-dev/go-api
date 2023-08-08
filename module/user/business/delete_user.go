package userbusiness

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
)

type DeleteUserStorage interface {
	Find(context context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Delete(context context.Context, id uuid.UUID) error
}

type deleteUserBusiness struct {
	storage DeleteUserStorage
}

func NewDeleteUserBusiness(storage DeleteUserStorage) *deleteUserBusiness {
	return &deleteUserBusiness{storage: storage}
}

func (business *deleteUserBusiness) DeleteUser(context context.Context, id uuid.UUID) error {
	conditions := map[string]interface{}{"id": id}

	_, err := business.storage.Find(context, conditions)

	if err == common.ErrRecordNotFound {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if err != nil {
		return common.ErrDB(err)
	}

	if err := business.storage.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	return nil
}
