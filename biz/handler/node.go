package handler

import (
	"net/http"

	"github.com/apusnetwork/open-api/biz/dal"
	"github.com/gin-gonic/gin"
)

func GetNodes(c *gin.Context) {
	var req = struct {
		onlyActive bool `form:"onlyActive" json:"onlyActive"`
	}{}

	nodes, err := dal.GetNodes(c, dal.GetDBConn(c), req.onlyActive)
	if err != nil {
		c.JSON(http.StatusBadRequest, c.Error(err))
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func NodeReport(c *gin.Context) {

}
