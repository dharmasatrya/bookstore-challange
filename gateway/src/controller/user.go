package controller

import (
	"fmt"
	"gateway/entity"
	"gateway/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

// CreateUser godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param order body entity.RegisterRequest true "Register input"
// @Success 201 {object} entity.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func (h *userController) RegisterUser(c echo.Context) error {
	var req entity.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
	}

	req.Password = string(hashedPassword)

	fmt.Println(req.Password)

	status, response := h.userService.RegisterUser(req)

	return c.JSON(status, response)
}

// LoginUser godoc
// @Summary Login
// @Tags users
// @Accept json
// @Produce json
// @Param order body entity.LoginRequest true "Login Information"
// @Success 201 {object} entity.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/login [post]
func (h *userController) LoginUser(c echo.Context) error {
	var req entity.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.userService.LoginUser(req)

	return c.JSON(status, response)
}
