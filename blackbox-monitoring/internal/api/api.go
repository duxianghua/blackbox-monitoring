package api

import (
	"github.com/duxianghua/blackbox-monitoring/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	router.Register(r)
	r.Run()
}
