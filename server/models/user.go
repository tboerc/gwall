package models

import (
	"github.com/tboerc/gwall/server/password"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string
	Password []byte
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	hash, err := password.Hash(u.Password)
	if err != nil {
		return
	}

	u.Password = hash

	return
}

func (u *User) Create() error {
	r := db.Create(u)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func init() {
	models = append(models, &User{})
}
