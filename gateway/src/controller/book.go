package controller

import (
	"gateway/entity"
	"gateway/src/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) *bookController {
	return &bookController{bookService}
}

// Create book godoc
// @Summary Register a new book
// @Tags books
// @Accept json
// @Produce json
// @Param order body entity.CreateBookInput true "book input"
// @Success 201 {object} entity.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *bookController) CreateBook(c echo.Context) error {
	var req entity.CreateBookInput

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.bookService.CreateBook(token, req)

	return c.JSON(status, response)
}
