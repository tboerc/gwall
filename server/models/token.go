package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID     string `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID string `type:"uuid"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	t.ID = id.String()

	return
}

func (t *Token) Create() error {
	q := db.Where("user_id = ?", t.UserID).Delete(&Token{})
	if q.Error != nil {
		return q.Error
	}

	q = db.Create(t)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func init() {
	models = append(models, &Token{})
}
