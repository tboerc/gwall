package models

import (
	"github.com/tboerc/gwall/server/messages"
	"github.com/tboerc/gwall/server/password"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username string `gorm:"unique"`
	Password []byte
	Tokens   []Token
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
	q := db.Create(u)
	if q.Error != nil {
		return q.Error
	}

	return nil
}

func (u *User) Login() ([]byte, error) {
	dbu := &User{}
	if q := db.First(dbu, "username = ?", u.Username); q.Error != nil {
		return nil, messages.ErrUserNotMatch
	}

	if m, _ := password.Compare(u.Password, dbu.Password); !m {
		return nil, messages.ErrUserNotMatch
	}

	t := &Token{UserID: dbu.ID}
	if err := t.Create(); err != nil {
		return nil, messages.ErrUserNotMatch
	}

	return []byte(t.ID), nil
}

func init() {
	models = append(models, &User{})
}
