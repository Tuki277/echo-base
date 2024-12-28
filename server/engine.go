package server

import (
	"gorm.io/gorm"
)

type EngineInterface interface {
	Serve()
	Address() string
}

type Engine struct {
	db      *gorm.DB
	address string
}
