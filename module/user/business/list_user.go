package userbusiness

import (
	"context"
	"go-api/common"
	usermodel "go-api/module/user/model"
)

type ListUserStorage interface {
	List(context context.Context, filter *usermodel.Filter, paging *common.Paging, moreInfo ...string) ([]usermodel.User, error)
}

type listUserBusiness struct {
	storage ListUserStorage
}

func NewListUserBusiness(storage ListUserStorage) *listUserBusiness {
	return &listUserBusiness{storage: storage}
}

func (business *listUserBusiness) ListUser(context context.Context, filter *usermodel.Filter, paging *common.Paging) ([]usermodel.User, error) {
	users, err := business.storage.List(context, filter, paging)

	for i := range users {
		users[i].Mask(false)
	}

	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return users, nil
}
