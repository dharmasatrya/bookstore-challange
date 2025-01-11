package controller

import (
	"fmt"
	"gateway/entity"
	"gateway/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type gatewayController struct {
	gatewayService service.GatewayService
}

func NewGatewayController(gatewayService service.GatewayService) *gatewayController {
	return &gatewayController{gatewayService}
}

// CreateUser godoc
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param order body entity.CreateUserInput true "Register input"
// @Success 201 {object} entity.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/register [post]
func (h *gatewayController) RegisterUser(c echo.Context) error {
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

	status, response := h.gatewayService.RegisterUser(req)

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
func (h *gatewayController) LoginUser(c echo.Context) error {
	var req entity.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.gatewayService.LoginUser(req)

	return c.JSON(status, response)
}
