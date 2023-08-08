package usermodel

import (
	"errors"
	"go-api/common"
	"net/mail"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	EntityName = "User"
	TableName  = "users"
)

type User struct {
	common.BaseModel `json:",inline"`
	FirstName        string        `json:"first_name" gorm:"column:first_name;type:varchar(128);not null"`
	LastName         string        `json:"last_name" gorm:"column:last_name;type:varchar(128);not null"`
	Avatar           *common.Image `json:"avatar" gorm:"column:avatar;type:jsonb;default:null"`
	Phone            string        `json:"phone" gorm:"column:phone;type:varchar(10);not null"`
	Email            string        `json:"email" gorm:"column:email;type:varchar(256);not null"`
	Password         string        `json:"-" gorm:"column:password;type:text;not null"`
	IsLocked         bool          `json:"is_locked" gorm:"column:is_locked;default:false"`
	VerificationCode *string       `json:"-" gorm:"column:verification_code;type:text;default:null"`
	Verified         bool          `json:"verified" gorm:"column:verified;default:false"`
	Token            *string       `json:"-" gorm:"column:token;default:null"`
	Role             string        `json:"role" gorm:"column:role;type:varchar(32);default:'user'"`
}

func (User) TableName() string { return TableName }

func (u *User) Mask(isAdmin bool) {
	u.FakeID = common.EncodeUID(u.ID)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}

type UserCreate struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u UserCreate) Validate() error {
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

type UserUpdate struct {
	FirstName       *string       `json:"first_name"`
	LastName        *string       `json:"last_name"`
	Avatar          *common.Image `json:"avatar"`
	Phone           *string       `json:"phone"`
	Password        *string       `json:"password"`
	ConfirmPassword *string       `json:"confirm_password"`
}

func (u UserUpdate) Validate() error {
	if u.FirstName != nil {
		*u.FirstName = strings.TrimSpace(*u.FirstName)
		if *u.FirstName == "" {
			return ErrFirstNameIsEmpty
		}

		if strings.ContainsAny(*u.FirstName, "!@#$%^&*()_+{}|:<>?0123456789") {
			return ErrFirstNameIsInvalid
		}
	}

	if u.LastName != nil {
		*u.LastName = strings.TrimSpace(*u.LastName)
		if *u.LastName == "" {
			return ErrLastNameIsEmpty
		}

		if strings.ContainsAny(*u.LastName, "!@#$%^&*()_+{}|:<>?0123456789") {
			return ErrLastNameIsInvalid
		}
	}

	if u.Phone != nil {
		*u.Phone = strings.TrimSpace(*u.Phone)
		if *u.Phone == "" {
			return ErrPhoneIsEmpty
		}

		if _, err := strconv.Atoi(*u.Phone); err != nil || len(*u.Phone) != 10 {
			return ErrPhoneIsInvalid
		}
	}

	if u.Password != nil {
		if !regexp.MustCompile(`^[^ ]{8,32}$`).MatchString(*u.Password) {
			return ErrPasswordIsInvalid
		}

		if u.ConfirmPassword == nil {
			return ErrConfirmPasswordIsEmpty
		}

		if *u.Password != *u.ConfirmPassword {
			return ErrPasswordNotMatch
		}
	}

	return nil
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
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

	ErrEmailExisted = common.NewCustomErrorResponse(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUserIsLocked = common.NewCustomErrorResponse(
		errors.New("user is locked"),
		"user is locked",
		"ErrUserIsLocked",
	)

	ErrEmailOrPasswordIsIncorrect = common.NewCustomErrorResponse(
		errors.New("email or password is incorrect"),
		"email or password is incorrect",
		"ErrEmailOrPasswordIsIncorrect",
	)
)
