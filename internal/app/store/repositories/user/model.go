package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID                int        `gorm:"primaryKey"`
	RoleID            int        `gorm:"column:role_id"`
	Email             string     `gorm:"column:email"`
	Phone             string     `gorm:"column:phone"`
	PhoneVerifiedAt   *time.Time `gorm:"column:phone_verified_at"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
	EmailVerifiedAt   *time.Time `gorm:"column:email_verified_at"`
	FirstName         string     `gorm:"column:first_name"`
	LastName          string     `gorm:"column:last_name"`
	About             string     `gorm:"column:about"`
	Age               int        `gorm:"age"`
	IDConfirmation    *time.Time `gorm:"column:id_confirmation"`
	Password          string     `gorm:"-"`
	EncryptedPassword string     `gorm:"column:password"`
}

// Validate
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Length(6, 100), validation.By(requiredIf(u.EncryptedPassword == ""))),
	)
}

// BeforeCreate
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}
	return nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
