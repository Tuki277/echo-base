package hash_password

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

const (
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 2002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 2003
	ERROR_AUTH_TOKEN               = 2004
	ERROR_PRIVATE_KEY              = 2005
	ERROR_PRIVATE_KEY_FAIL         = 2006
	ERROR_USER_IS_BLOCK            = 2007
	ERROR_USER_IS_RESTRICTED       = 2008
)
const (
	Salt   = "REC"
	JWTkey = "REC"
)

var MsgFlags = map[int]string{
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token is invalid",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Your session has timed out. Please login again.",
	ERROR_AUTH_TOKEN:               "Token is error, please try again",
	ERROR_PRIVATE_KEY:              "Private key is error, please try again",
	ERROR_PRIVATE_KEY_FAIL:         "Private key is invalid",
	ERROR_USER_IS_BLOCK:            "User is block",
	ERROR_USER_IS_RESTRICTED:       "User is restricted",
}

func GetMsg(code int) string {
	msg, _ := MsgFlags[code]
	return msg
}

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Status   uint   `json:"status"`
	Role     uint   `json:"role"`
	jwt.StandardClaims
}

func NewJWT(JwtSecret string) *JWT {
	return &JWT{
		[]byte(JwtSecret),
	}
}

func (j *JWT) GenerateToken(Id uint, FullName string, Email string, Status uint, Role uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(15 * 24 * time.Hour)

	claims := CustomClaims{
		Id,
		FullName,
		Email,
		Status,
		Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "REC",
		},
	}
	return j.CreateToken(claims)
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
