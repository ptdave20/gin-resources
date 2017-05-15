package auth

import (
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	HAS_TOKEN = "has_token"
	TOKEN_ERR = "token_err"
	AUTHENTICATED = "authenticated"
)

type (
	TokenVerification func(*gin.Context,string) (bool,error)
)

func BearerAuthentication(verification TokenVerification, checkCookie bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth:=c.Request.Header.Get("Authentication")
		c.Set(HAS_TOKEN, len(auth) == 0)

		a,e:=verification(c,auth);
		if e!=nil {
			c.Set(TOKEN_ERR,e)
			c.Set(AUTHENTICATED, false)
			c.Next()
		} else if a {
			c.Set(AUTHENTICATED, true)
		}
	}
}
