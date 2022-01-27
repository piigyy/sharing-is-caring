package model

import (
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
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
