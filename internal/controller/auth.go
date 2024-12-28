package controller

import (
	"echo-base/internal/contract"
	"echo-base/internal/repository"
	"echo-base/internal/service"
	"echo-base/utils/constants"
	"echo-base/utils/responses"
	"echo-base/utils/validation"
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type AuthController struct {
	server *echo.Echo
	db     *gorm.DB
}

func NewAuthController(db *gorm.DB, server *echo.Echo) *AuthController {
	return &AuthController{db: db, server: server}
}

func (c *AuthController) RegisterHandler() {
	group := c.server.Group("/admin/auth")

	group.POST("/login", c.Login)
	group.POST("/register", c.Register)
}

func (c *AuthController) Login(e echo.Context) error {
	request := new(contract.LoginRequest)
	if err := e.Bind(request); err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	isValid, err := validation.CommonlyValidate(request)
	if !isValid {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	result, err := c.Service().Login(request)
	if err != nil && err != gorm.ErrRecordNotFound {
		data := responses.ResponseData(result, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	} else if err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	data := responses.ResponseData(result, nil, err, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}

func (c *AuthController) Register(e echo.Context) error {
	request := new(contract.RegisterRequest)
	if err := e.Bind(request); err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	isValid, err := validation.CommonlyValidate(request)
	if !isValid {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	result, err := c.Service().Register(request)
	if err != nil && err.Error() != constants.ErrEmailExisting {
		data := responses.ResponseData(result, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	} else if err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	data := responses.ResponseData(result, nil, err, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}

func (c *AuthController) Service() *service.AuthService {
	return service.NewAuthService(repository.NewUserRepository(c.db))
}
