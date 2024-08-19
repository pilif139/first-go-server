package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func (uh *UserHandler) CreateUser(c echo.Context) error {
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	if err := uh.DB.Create(&u).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, u)
}

func (uh *UserHandler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, uh.DB.First(&User{}, id))
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	u.ID = uint(id)
	if err := uh.DB.Save(u).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, u)
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	u := uh.DB.First(&User{}, id)
	if err := uh.DB.Delete(u).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (uh *UserHandler) GetAllUsers(c echo.Context) error {
	var users []User
	result := uh.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, users)
}
