package gincasbin

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// Middleware returns the HandlerFunc, uses a Casbin enforcer as input
func Middleware(e *casbin.Enforcer, soaFn func(c *gin.Context) (string, string, string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		subject, object, action := soaFn(c)
		allowed, err := e.Enforce(subject, object, action)
		if err != nil {
			panic(err)
		}

		if !allowed {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
