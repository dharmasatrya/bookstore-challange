package controller

import (
	"gateway/entity"
	"gateway/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type borrowController struct {
	borrowService service.BorrowService
}

func NewBorrowController(borrowService service.BorrowService) *borrowController {
	return &borrowController{borrowService}
}

// @Summary Borrow a book
// @Tags borrows
// @Accept json
// @Produce json
// @Param order body entity.BorrowBookInput true "borrow input"
// @Success 201 {object} entity.BorrowedBook
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /borrows [post]
func (h *borrowController) BorrowBook(c echo.Context) error {
	var req entity.BorrowBookInput

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.borrowService.BorrowBook(token, req)
	return c.JSON(status, response)
}

// @Summary Edit a borrowed book
// @Tags borrows
// @Accept json
// @Produce json
// @Param order body entity.EditBorrowRequest true "edit borrow input"
// @Success 200 {object} entity.BorrowedBook
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /borrows/:id [put]
func (h *borrowController) EditBorrowedBook(c echo.Context) error {
	var req entity.EditBorrowRequest

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	borrowId := c.Param("id")

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.borrowService.EditBorrowedBook(token, borrowId, req)
	return c.JSON(status, response)
}
