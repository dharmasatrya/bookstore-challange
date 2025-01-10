package controller

import (
	"gateway/entity"
	"gateway/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type gatewayController struct {
	gatewayService service.GatewayService
}

func NewGatewayController(gatewayService service.GatewayService) *gatewayController {
	return &gatewayController{gatewayService}
}

// CreateOrder godoc
// @Summary Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body entity.CreateOrderInput true "Order Information"
// @Success 201 {object} entity.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (h *gatewayController) RegisterUser(c echo.Context) error {
	var req entity.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.gatewayService.RegisterUser(req)

	return c.JSON(status, response)
}
