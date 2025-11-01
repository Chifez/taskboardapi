package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID 			uint 			`gorm:"primarykey" json:"id"`
	Email		string			`gorm:"unique;not null" json:"email"`
	Password	string			`gorm:"not null" json:"-"`
	Name		string			`gorm:"not null" json:"name"`
	Projects	[]Project		`gorm:"foreignKey:UserID" json:"projects,omitempty"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index" json:"-"`
}

func (u *User) HashPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

}