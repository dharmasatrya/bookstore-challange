package routes

import (
	"gateway/src/controller"
	"gateway/src/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"google.golang.org/grpc"
)

func NewRouter() *echo.Echo {

	userConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did'nt connect : %v", err)
	}

	gatewayService := service.NewGatewayService(userConnection)

	gatewayController := controller.NewGatewayController(gatewayService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	u := e.Group("/users")
	u.POST("/register", gatewayController.RegisterUser)

	return e
}
