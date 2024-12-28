package server

import (
	"echo-base/internal/controller"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	swagger "github.com/swaggo/echo-swagger"
)

type EchoEngine struct {
	Engine
}

func NewEchoEngine(db *gorm.DB, address string) *EchoEngine {
	return &EchoEngine{Engine{db: db, address: address}}
}

func (s *EchoEngine) Serve() {
	server := echo.New()
	server.Use(middleware.CORS())

	server.GET("/swagger/*", swagger.WrapHandler)

	controller.NewUserController(s.db, server).RegisterHandler()
	controller.NewHealthCheckController(server).RegisterHandler()
	controller.NewAuthController(s.db, server).RegisterHandler()

	server.Logger.Fatal(server.Start(s.address))
}

func (s *EchoEngine) Address() string {
	return s.address
}
