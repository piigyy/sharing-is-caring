package model

import (
	"net/mail"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Phone     string             `bson:"phone" json:"phone"`
	Password  string             `bson:"password" json:"-"`
	CreateAt  time.Time          `bson:"createdAt" json:"createAt"`
	Updated   time.Time          `bson:"updatedAt" json:"updatedAt"`
	DeletedAt *time.Time         `bson:"deletedAt" json:"-"`
}

func (u *User) SetPhoneNumber(phone string) {
	u.Phone = strings.Join(strings.Fields(phone), "")
}

func (u *User) ValidateEmail() bool {
	_, err := mail.ParseAddress(u.Email)
	return err == nil
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}
