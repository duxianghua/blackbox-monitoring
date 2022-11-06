package router

import (
	"github.com/duxianghua/blackbox-monitoring/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func BlackboxApi(r *gin.Engine) {
	r.GET("/get", (&handler.BlackboxExporter{}).Get)
	r.GET("/gethttpprobe", (&handler.BlackboxExporter{}).GetHttpProbe)
}
