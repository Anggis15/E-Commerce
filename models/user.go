package models

import (
	"FP-BDS-Sanbercode-Go-50-anggi/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	// User
	User struct {
		Id         int `gorm:"primary_key" json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		Gender       string `json:"gender"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		IsActive int `json:"isActive"`
		Products []Products `josn:"-"`
	}
)

func VerfyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck (email string, password string, db *gorm.DB) (string, error){
	var err error

	u := User{}

	err = db.Model(User{}).Where("email = ?", email).Take(&u).Error

	err = VerfyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		return "", nil
	}

	token, err := token.GenerateToken(uint(u.Id))

	if err != nil{
		return "", err
	}

	return token, nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
    //turn password into hash
    hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if errPassword != nil {
        return &User{}, errPassword
    }
    u.Password = string(hashedPassword)
    //remove spaces in username
    u.Email = html.EscapeString(strings.TrimSpace(u.Email))

    var err error = db.Create(&u).Error
    if err != nil {
        return &User{}, err
    }
    return u, nil
}