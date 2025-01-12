package routes

import (
	"gateway/src/controller"
	"gateway/src/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRouter() *echo.Echo {

	// userConnection, err := grpc.Dial("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	userConnection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did'nt connect : %v", err)
	}

	bookConnection, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did'nt connect : %v", err)
	}

	userService := service.NewUserService(userConnection)
	userController := controller.NewUserController(userService)

	bookService := service.NewBookService(bookConnection)
	bookController := controller.NewBookController(bookService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	u := e.Group("/users")
	u.POST("/register", userController.RegisterUser)
	u.POST("/login", userController.LoginUser)

	b := e.Group("/books")
	b.POST("", bookController.CreateBook)
	b.PUT("/:id", bookController.EditBook)
	b.DELETE("/:id", bookController.DeleteBook)

	return e
}
