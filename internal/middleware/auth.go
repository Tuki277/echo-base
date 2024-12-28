package middleware

import (
	"echo-base/internal/service"
	"echo-base/utils/hash_password"
	"echo-base/utils/responses"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		token := e.Request().Header.Get("Authorization")
		if token == "" {
			data := responses.ResponseData(nil, nil, errors.New(hash_password.GetMsg(hash_password.ERROR_AUTH_TOKEN)), http.StatusBadRequest)
			return e.JSONPretty(http.StatusBadRequest, data, "  ")
		}
		j := hash_password.NewJWT(service.JWTkey)
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == hash_password.TokenExpired {
				data := responses.ResponseData(nil, nil, errors.New(hash_password.GetMsg(hash_password.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)), http.StatusBadRequest)
				return e.JSONPretty(http.StatusBadRequest, data, "  ")
			}
			data := responses.ResponseData(nil, nil, errors.New(hash_password.GetMsg(hash_password.ERROR_AUTH_CHECK_TOKEN_FAIL)), http.StatusBadRequest)
			return e.JSONPretty(http.StatusBadRequest, data, "  ")
		}
		e.Set("claims", claims)
		err = next(e)
		return err
	}
}
