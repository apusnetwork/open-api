package router

import (
	"log"

	"github.com/apusnetwork/open-api/biz/handler"
	"github.com/apusnetwork/open-api/biz/middleware/gincasbin"
	"github.com/apusnetwork/open-api/biz/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	r.GET("/ping", handler.Ping)

	r.POST("/apikey", handler.CreateApikey)
	r.GET("/apikey", handler.GetApikey)

	authorized := r.Group("/")

	e, _ := casbin.NewEnforcer("./config/model.conf", "./config/policy.csv")
	authorized.Use(gincasbin.Middleware(e, soaFromRequest))
	{
		authorized.POST("/apikey/check")

		authorized.GET("/health/xxx", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		authorized.GET("/nodes", handler.GetNodes)
		authorized.POST("/node/report", handler.NodeReport)
	}

}

func soaFromRequest(c *gin.Context) (string, string, string) {
	subject := "anonymous"

	ak := c.GetHeader("accessKey")
	signature := c.GetHeader("signature")
	authz, err := service.CheckSecretKey(c, ak, "", signature)
	if err == nil {
		subject = authz.Role
	}

	object := c.Request.URL.Path
	action := c.Request.Method
	log.Println(subject, object, action)

	return subject, object, action
}
