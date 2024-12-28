package controller

import (
	"echo-base/internal/contract"
	"echo-base/internal/middleware"
	"echo-base/internal/repository"
	"echo-base/internal/service"
	"echo-base/utils/constants"
	"echo-base/utils/hash_password"
	"echo-base/utils/responses"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	server *echo.Echo
	db     *gorm.DB
}

func NewUserController(db *gorm.DB, server *echo.Echo) *UserController {
	return &UserController{db: db, server: server}
}
func (c *UserController) RegisterHandler() {
	group := c.server.Group("/admin/users", middleware.JWTAuth)

	//group.GET("/:id", c.Get)
	group.GET("", c.List)
	group.GET("/me", c.Me)

	group.PUT("/status", c.ChangeStatus)
	//group.PUT("/priority", c.UpdatePriority)
	group.DELETE("", c.Delete)
}
func (c *UserController) Delete(e echo.Context) error {
	claims := e.Get("claims")
	val, ok := claims.(*hash_password.CustomClaims)
	if !ok {
		data := responses.ResponseData(nil, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	}

	request := new(contract.DeleteUserReq)
	err := e.Bind(request)
	if err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	for _, v := range request.Ids {
		if v == val.ID {
			data := responses.ResponseData(nil, nil, errors.New("Cannot delete admin"), http.StatusBadRequest)
			return e.JSONPretty(http.StatusBadRequest, data, "  ")
		}
	}

	err = c.Service().Delete(request.Ids)
	if err != nil && err != gorm.ErrRecordNotFound {
		data := responses.ResponseData(nil, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	} else if err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}

	data := responses.ResponseData(nil, nil, err, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}
func (c *UserController) Me(e echo.Context) error {
	claims := e.Get("claims")
	val, ok := claims.(*hash_password.CustomClaims)
	if !ok {
		data := responses.ResponseData(nil, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	}
	result := &contract.UserResponse{
		Id:       val.ID,
		FullName: val.FullName,
		Email:    val.Email,
	}
	data := responses.ResponseData(result, nil, nil, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}
func (c *UserController) List(e echo.Context) error {
	result, err := c.Service().List()
	if err != nil {
		data := responses.ResponseData(result, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	}

	data := responses.ResponseData(result, nil, err, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}

func (c *UserController) ChangeStatus(e echo.Context) error {
	claims := e.Get("claims")
	val, ok := claims.(*hash_password.CustomClaims)
	if !ok {
		data := responses.ResponseData(nil, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	}

	request := new(contract.ChangeStatusRes)
	err := e.Bind(request)
	if err != nil {
		data := responses.ResponseData(nil, nil, err, http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	if request.Id == val.ID {
		data := responses.ResponseData(nil, nil, errors.New("Cannot change status of admin"), http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}
	err = c.Service().ChangeStatus(request)
	if err != nil && err != gorm.ErrRecordNotFound {
		data := responses.ResponseData(nil, nil, errors.New(constants.ErrInternalError), http.StatusInternalServerError)
		return e.JSONPretty(http.StatusInternalServerError, data, "  ")
	} else if err != nil {
		data := responses.ResponseData(nil, nil, errors.New(constants.UserNotExisting), http.StatusBadRequest)
		return e.JSONPretty(http.StatusBadRequest, data, "  ")
	}

	data := responses.ResponseData(nil, nil, err, http.StatusOK)
	return e.JSONPretty(http.StatusOK, data, "  ")
}

func (c *UserController) Service() *service.UserService {
	return service.NewUserService(repository.NewUserRepository(c.db))
}
