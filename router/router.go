package router

import (
	"github.com/gin-gonic/gin"
	"promethues_alertmanager_webhook/api/v1/dingding"
)

// Router 总路由/*
func Router(r *gin.RouterGroup) {
	chat := r.Group("/chat")

	// 模块分组
	dingding.Router(chat)
}
