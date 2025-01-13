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
	// userConnection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did'nt connect : %v", err)
	// }

	// bookConnection, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did'nt connect : %v", err)
	// }

	// borrowConnection, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did'nt connect : %v", err)
	// }

	userConnection, err := grpc.Dial("user-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to user service: %v", err)
	}

	bookConnection, err := grpc.Dial("book-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to book service: %v", err)
	}

	borrowConnection, err := grpc.Dial("borrow-service:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Didn't connect to borrow service: %v", err)
	}

	userService := service.NewUserService(userConnection)
	userController := controller.NewUserController(userService)

	bookService := service.NewBookService(bookConnection)
	bookController := controller.NewBookController(bookService)

	borrowService := service.NewBorrowService(borrowConnection)
	borrowController := controller.NewBorrowController(borrowService)

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
	b.GET("", bookController.GetAllBook)
	b.GET("/:id", bookController.GetBookById)

	br := e.Group("/borrow")
	br.POST("", borrowController.BorrowBook)
	br.PUT("/:id", borrowController.EditBorrowedBook)

	return e
}
