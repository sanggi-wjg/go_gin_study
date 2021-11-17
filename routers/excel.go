package routers

import (
	"go_gin_study/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Excel(c *gin.Context) {
	service.CreateExcel()
	c.String(http.StatusOK, "Users success")
}
