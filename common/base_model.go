package common

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `json:"-" gorm:"column:id;type:uuid;primary_key;default:gen_random_uuid()"`
	FakeID    string    `json:"id" gorm:"-"`
	Status    int       `json:"-" gorm:"column:status;type:smallint;default:1;index:idx_user_status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;autoUpdateTime"`
}
