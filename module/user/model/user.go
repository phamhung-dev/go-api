package usermodel

import (
	"errors"
	"go-api/common"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	common.BaseModel `json:",inline"`
	FirstName        string  `json:"first_name" gorm:"column:first_name;type:varchar(128);not null"`
	LastName         string  `json:"last_name" gorm:"column:last_name;type:varchar(128);not null"`
	Avatar           *string `json:"avatar" gorm:"column:avatar;default:null"`
	Phone            string  `json:"phone" gorm:"column:phone;type:varchar(10);not null"`
	Email            string  `json:"email" gorm:"column:email;type:varchar(256);not null"`
	Password         string  `json:"-" gorm:"column:password;not null"`
	IsLocked         bool    `json:"is_locked" gorm:"column:is_locked;default:false"`
	VerificationCode *string `json:"-" gorm:"column:verification_code;default:null"`
	Verified         bool    `json:"verified" gorm:"column:verified;default:false"`
	Token            *string `json:"-" gorm:"column:token;default:null"`
}

func (User) TableName() string { return "users" }

type UserCreateRequest struct {
	ID              uuid.UUID `json:"-" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()"`
	FirstName       string    `json:"first_name" binding:"required" gorm:"column:first_name"`
	LastName        string    `json:"last_name" binding:"required" gorm:"column:last_name"`
	Email           string    `json:"email" binding:"required" gorm:"column:email"`
	Phone           string    `json:"phone" binding:"required" gorm:"column:phone"`
	Password        string    `json:"password" binding:"required" gorm:"column:password"`
	ConfirmPassword string    `json:"confirm_password" binding:"required" gorm:"-:all"`
	CreatedAt       time.Time `json:"-" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `json:"-" gorm:"column:updated_at;autoUpdateTime"`
}

func (UserCreateRequest) TableName() string { return User{}.TableName() }

func (u *UserCreateRequest) Validate() error {
	u.FirstName = strings.TrimSpace(u.FirstName)

	if u.FirstName == "" {
		return ErrFirstNameIsEmpty
	}

	if strings.ContainsAny(u.FirstName, "!@#$%^&*()_+{}|:<>?0123456789") {
		return ErrFirstNameIsInvalid
	}

	u.LastName = strings.TrimSpace(u.LastName)

	if u.LastName == "" {
		return ErrLastNameIsEmpty
	}

	if strings.ContainsAny(u.LastName, "!@#$%^&*()_+{}|:<>?0123456789") {
		return ErrLastNameIsInvalid
	}

	u.Email = strings.TrimSpace(u.Email)

	if u.Email == "" {
		return ErrEmailIsEmpty
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return ErrEmailIsInvalid
	}

	u.Phone = strings.TrimSpace(u.Phone)

	if u.Phone == "" {
		return ErrPhoneIsEmpty
	}

	if _, err := strconv.Atoi(u.Phone); err != nil || len(u.Phone) != 10 {
		return ErrPhoneIsInvalid
	}

	if u.Password == "" {
		return ErrPasswordIsEmpty
	}

	if !regexp.MustCompile(`^[^ ]{8,32}$`).MatchString(u.Password) {
		return ErrPasswordIsInvalid
	}

	if u.ConfirmPassword == "" {
		return ErrConfirmPasswordIsEmpty
	}

	if u.Password != u.ConfirmPassword {
		return ErrPasswordNotMatch
	}

	return nil
}

func (u *UserCreateRequest) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

type UserCreateResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdateRequest struct {
	FirstName       *string   `json:"first_name" gorm:"column:first_name"`
	LastName        *string   `json:"last_name" gorm:"column:last_name"`
	Phone           *string   `json:"phone" gorm:"column:phone"`
	Password        *string   `json:"password" gorm:"column:password"`
	ConfirmPassword *string   `json:"confirm_password" gorm:"-:all"`
	UpdatedAt       time.Time `json:"-" gorm:"column:updated_at;autoUpdateTime"`
}

func (UserUpdateRequest) TableName() string { return User{}.TableName() }

type UserUpdateResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	UpdatedAt string `json:"updated_at"`
}

var (
	ErrFirstNameIsEmpty       = errors.New("first name is empty")
	ErrLastNameIsEmpty        = errors.New("last name is empty")
	ErrEmailIsEmpty           = errors.New("email is empty")
	ErrPhoneIsEmpty           = errors.New("phone is empty")
	ErrPasswordIsEmpty        = errors.New("password is empty")
	ErrConfirmPasswordIsEmpty = errors.New("confirm password is empty")
	ErrPasswordNotMatch       = errors.New("password not match")

	ErrFirstNameIsInvalid = errors.New("first name is invalid")
	ErrLastNameIsInvalid  = errors.New("last name is invalid")
	ErrEmailIsInvalid     = errors.New("email is invalid")
	ErrPhoneIsInvalid     = errors.New("phone is invalid")
	ErrPasswordIsInvalid  = errors.New("password is invalid")
)
