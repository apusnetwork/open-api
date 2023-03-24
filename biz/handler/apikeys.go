package handler

import (
	"net/http"

	"github.com/apusnetwork/open-api/biz/dal"
	"github.com/apusnetwork/open-api/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetApikey(c *gin.Context) {
	var req = struct {
		Address   string `form:"address" json:"address" binding:"required"`
		Signature string `form:"signature" json:"signature" binding:"required"`
	}{}

	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	// web3签名验证

	authz, err := dal.GetAuthzByAccessKey(c, dal.GetDBConn(c), req.Address)
	if err != nil {
		// hlog.CtxErrorf(ctx, "method `GetApikey` failed, detail: %s", err)
		c.JSON(http.StatusInternalServerError, c.Error(err))
		return
	}
	if authz == nil {
		if err != nil {
			// hlog.CtxErrorf(ctx, "method `GetApikey` failed, detail: %s", err)
			c.JSON(http.StatusInternalServerError, c.Error(err))
			return
		}
	}
	c.JSON(http.StatusOK, authz)
}

func CreateApikey(c *gin.Context) {
	var req = struct {
		Address   string `form:"address" json:"address" binding:"required"`
		Signature string `form:"signature" json:"signature" binding:"required"`
		Role      string `form:"role" json:"role" binding:"required,oneof=miner developer"`
	}{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	// web3签名验证

	// 生成secert key
	secretKey := uuid.New().String()
	authz, err := dal.CreateOrUpdate(c, dal.GetDBConn(c), req.Address, secretKey, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	c.JSON(http.StatusOK, authz)
}

func CheckApikey(c *gin.Context) {
	var req = struct {
		AccessKey string `form:"address" json:"address" binding:"required"`
		Message   string `form:"message" json:"message" binding:"required"`
		Signature string `form:"signature" json:"signature" binding:"required"`
	}{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}

	_, err = service.CheckSecretKey(c, req.AccessKey, "", req.Signature)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	c.JSON(http.StatusOK, "ok")
}
