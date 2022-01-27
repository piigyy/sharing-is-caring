package model_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/piigyy/sharing-is-caring/internal/auth/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUser_SetPhoneNumber(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		Email     string
		Phone     string
		Password  string
		CreateAt  time.Time
		Updated   time.Time
		DeletedAt *time.Time
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "trim whitespaces",
			fields: fields{
				ID:        primitive.NewObjectID(),
				Email:     faker.Email(),
				Password:  faker.Password(),
				CreateAt:  time.Now(),
				Updated:   time.Now(),
				DeletedAt: nil,
			},
			args: args{
				phone: "62 896 5887 6666",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &model.User{
				ID:        tt.fields.ID,
				Email:     tt.fields.Email,
				Phone:     tt.fields.Phone,
				Password:  tt.fields.Password,
				CreateAt:  tt.fields.CreateAt,
				Updated:   tt.fields.Updated,
				DeletedAt: tt.fields.DeletedAt,
			}

			u.SetPhoneNumber(tt.args.phone)
			assert.Equal(t, "6289658876666", u.Phone)
		})
	}
}

func TestUser_ValidateEmail(t *testing.T) {
	type fields struct {
		ID        primitive.ObjectID
		Name      string
		Email     string
		Phone     string
		Password  string
		CreateAt  time.Time
		Updated   time.Time
		DeletedAt *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "email address valid",
			fields: fields{
				Email: faker.Email(),
			},
			want: true,
		},
		{
			name: "email address empty",
			fields: fields{
				Email: "",
			},
			want: false,
		},
		{
			name: "email address invalid",
			fields: fields{
				Email: "hello world",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &model.User{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Email:     tt.fields.Email,
				Phone:     tt.fields.Phone,
				Password:  tt.fields.Password,
				CreateAt:  tt.fields.CreateAt,
				Updated:   tt.fields.Updated,
				DeletedAt: tt.fields.DeletedAt,
			}
			actual := u.ValidateEmail()
			assert.Equal(t, tt.want, actual, fmt.Sprintf("User.ValidateEmail() = %v, want %v", actual, tt.want))
		})
	}
}
