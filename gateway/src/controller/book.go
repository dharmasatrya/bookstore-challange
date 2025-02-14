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

// Edit book godoc
// @Summary Edit a book
// @Tags books
// @Accept json
// @Produce json
// @Param order body entity.EditBookRequest true "book input"
// @Success 201 {object} entity.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/:id [put]
func (h *bookController) EditBook(c echo.Context) error {
	var req entity.EditBookRequest

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	bookId := c.Param("id")

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	status, response := h.bookService.EditBook(token, bookId, req)

	return c.JSON(status, response)
}

// Delete book godoc
// @Summary delete a book
// @Tags books
// @Accept json
// @Produce json
// @Param order body entity.EditBookRequest true "book input"
// @Success 201 {object} entity.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/:id [delete]
func (h *bookController) DeleteBook(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	bookId := c.Param("id")

	status, response := h.bookService.DeleteBook(token, bookId)

	return c.JSON(status, response)
}

// GetBookById godoc
// @Summary Get a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} entity.Book
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [get]
func (h *bookController) GetBookById(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	bookId := c.Param("id")

	status, response := h.bookService.GetBookById(token, bookId)

	return c.JSON(status, response)
}

// GetAllBook godoc
// @Summary Get all books
// @Description Retrieve a list of all available books
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entity.Book
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *bookController) GetAllBook(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "No token provided",
		})
	}

	status, response := h.bookService.GetAllBooks(token)

	return c.JSON(status, response)
}
