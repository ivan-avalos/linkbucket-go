/* 
 *  user-controller.go
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

package controllers

import (
	"net/http"

	"github.com/ivan-avalos/linkbucket/database"
	"github.com/ivan-avalos/linkbucket/utils"
	"github.com/labstack/echo"
)

// CreateUser creates account for user
func CreateUser(c echo.Context) (err error) {
	registerUser := new(database.RegisterUser)
	if err = c.Bind(registerUser); err != nil {
		return utils.ProcessError(err)
	}
	if err = c.Validate(registerUser); err != nil {
		return utils.ProcessError(err)
	}
	user := registerUser.GetUser()
	if err = user.Create(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, user.GetResponseUser())
}

// Authenticate logs user in
func Authenticate(c echo.Context) (err error) {
	loginUser := new(database.LoginUser)
	if err = c.Bind(loginUser); err != nil {
		return utils.ProcessError(err)
	}
	if err = c.Validate(loginUser); err != nil {
		return utils.ProcessError(err)
	}
	user, err := database.Authenticate(loginUser.Email, loginUser.Password)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, user.GetResponseUser())
}

// GetUser retrieves user from DB
func GetUser(c echo.Context) (err error) {
	userID := utils.GetJWTUserID(c)
	user, err := database.GetUser(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, user.GetResponseUser())
}

// UpdateUser modifies user from DB
func UpdateUser(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	user, err := database.GetUser(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	updateUser := new(database.UpdateUser)
	if err := c.Bind(updateUser); err != nil {
		return utils.ProcessError(err)
	}
	if err := c.Validate(updateUser); err != nil {
		return utils.ProcessError(err)
	}
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		if err := database.FieldIsSameOrUnique("users", userID, "email", updateUser.Email); err != nil {
			return err
		}
		user.Email = updateUser.Email
	}
	if err := user.Update(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, user.GetResponseUser())
}

// DeleteUser removes user from DB
func DeleteUser(c echo.Context) (err error) {
	userID := utils.GetJWTUserID(c)
	user, err := database.GetUser(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	if err = user.Delete(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, user.GetResponseUser())
}
