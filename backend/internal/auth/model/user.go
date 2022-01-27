package model

import (
	"net/mail"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
	Password  string             `bson:"password"`
	CreateAt  time.Time          `bson:"createdAt"`
	Updated   time.Time          `bson:"updatedAt"`
	DeletedAt *time.Time         `bson:"deletedAt"`
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
