package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

const (
	HAS_TOKEN = "has_token"
	TOKEN_ERR = "token_err"
	AUTHENTICATED = "authenticated"

)

type (
	TokenVerification func(*gin.Context,string) (bool,error)
)

var (
	ERROR_NOTOKEN = errors.New("Request lacks a token")
	ERROR_FAILEDVERIFICATION = errors.New("Token failed verification")
)

func getCookieToken(c *gin.Context) (string) {
	if v,err:=c.Cookie("authentication"); err!=nil {
		return ""
	} else {
		return v
	}
}

func getBearerToken(c *gin.Context) string {
	return c.Request.Header.Get("Authentication")
}

func BearerAuthentication(verification TokenVerification, checkCookie bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth:=getBearerToken(c)

		if len(auth) == 0 && checkCookie {
			auth=getCookieToken(c)
		}

		if len(auth) == 0 {
			c.Set(HAS_TOKEN, false)
			c.Set(TOKEN_ERR, ERROR_NOTOKEN)
		} else {
			auth_pass,e:=verification(c,auth);
			if e!=nil {
				c.Set(TOKEN_ERR,e)
			}
			c.Set(AUTHENTICATED, auth_pass)

		}

		c.Next()
	}
}

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.MustGet(AUTHENTICATED).(bool) {
			// We are authenticated
			c.Next()
		} else {
			c.IndentedJSON(http.StatusUnauthorized,c.MustGet(TOKEN_ERR).(*error))
			c.Abort()
		}
	}
}