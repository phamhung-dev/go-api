package userbusiness

import (
	"context"
	"errors"
	usermodel "go-api/module/user/model"
)

type CreateUserStore interface {
	FindDataWithConditions(context context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
	Create(context context.Context, data *usermodel.UserCreateRequest) error
}

type createUserBusiness struct {
	store CreateUserStore
}

func NewCreateUserBusiness(store CreateUserStore) *createUserBusiness {
	return &createUserBusiness{store: store}
}

func (business *createUserBusiness) CreateUser(context context.Context, data *usermodel.UserCreateRequest) (*usermodel.UserCreateResponse, error) {
	conditions := map[string]interface{}{
		"email": data.Email,
	}

	if err := data.Validate(); err != nil {
		return nil, err
	}

	if user, err := business.store.FindDataWithConditions(context, conditions); err == nil || user != nil {
		return nil, ErrEmailAlreadyExists
	}

	if err := business.store.Create(context, data); err != nil {
		return nil, err
	}

	response := &usermodel.UserCreateResponse{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Phone:     data.Phone,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
	}

	return response, nil
}

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)
