package userbusiness

import (
	"context"
	"errors"
	usermodel "go-api/module/user/model"

	"github.com/google/uuid"
)

type DeleteUserStore interface {
	FindDataWithConditions(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	Delete(context context.Context, id uuid.UUID) error
}

type deleteUserBusiness struct {
	store DeleteUserStore
}

func NewDeleteUserBusiness(store DeleteUserStore) *deleteUserBusiness {
	return &deleteUserBusiness{store: store}
}

func (business *deleteUserBusiness) DeleteUser(context context.Context, id uuid.UUID) error {
	conditions := map[string]interface{}{"id": id}
	if user, err := business.store.FindDataWithConditions(context, conditions); err != nil || user == nil {
		return ErrUserNotFound
	}

	if err := business.store.Delete(context, id); err != nil {
		return err
	}

	return nil
}

var (
	ErrUserNotFound = errors.New("user not found")
)
