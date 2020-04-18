/*
 *  user.go
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package database

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ivan-avalos/linkbucket/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	// User represents a user
	User struct {
		ID        uint `gorm:"primary_key"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`

		Name          string `gorm:"not null"`
		Email         string `gorm:"type:varchar(100);unique_index"`
		Password      string `gorm:"not null"`
		RememberToken string
		Token         string `gorm:"-"`
		Tags          []Tag
	}

	// LoginUser represents user
	LoginUser struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// RegisterUser represents user
	RegisterUser struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email,unique=users.email"`
		Password string `json:"password" validate:"required,gt=8"`
	}

	// UpdateUser represents user
	UpdateUser struct {
		Name  string `json:"name"`
		Email string `json:"email" validate:"omitempty,email"`
	}

	// ResponseUser represents a response version of User
	ResponseUser struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// GetUser returns User from RegisterUser
func (registerUser *RegisterUser) GetUser() *User {
	return &User{
		Name:     registerUser.Name,
		Email:    registerUser.Email,
		Password: registerUser.Password,
	}
}

// GetUser returns User from UpdateUser
func (updateUser *UpdateUser) GetUser() *User {
	return &User{
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}
}

// GetResponseUser returns response version of User
func (user *User) GetResponseUser() *ResponseUser {
	respUser := &ResponseUser{}
	respUser.ID = user.ID
	respUser.Name = user.Name
	respUser.Email = user.Email
	respUser.Token = user.Token
	respUser.CreatedAt = user.CreatedAt
	respUser.UpdatedAt = user.UpdatedAt
	return respUser
}

// GenerateJWT generates a JWT token for a user
func (user *User) GenerateJWT() (string, error) {
	tk := &utils.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	return token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))
}

// Authenticate validates user data and generates JWT token
func Authenticate(email, password string) (*User, error) {
	user := &User{}
	err := DB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	user.Password = ""
	tokenString, _ := user.GenerateJWT()
	user.Token = tokenString
	return user, nil
}

// Create inserts a user into DB
func (user *User) Create() error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := DB().Create(user).Error; err != nil {
		return err
	}

	token, err := user.GenerateJWT()
	if err != nil {
		return err
	}

	user.Token = token
	user.Password = ""
	return nil
}

// GetUser retrieves a user from DB
func GetUser(id uint) (*User, error) {
	user := &User{}
	err := DB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

// Update modifies a user in DB
func (user *User) Update() error {
	return DB().Save(user).Error
}

// Delete removes a user from DB
func (user *User) Delete() error {
	return DB().Delete(user).Error
}
