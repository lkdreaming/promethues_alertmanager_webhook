package dingding

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.POST("/v1/dingTalk/sendMarkdown", PostMarkdown)
}
