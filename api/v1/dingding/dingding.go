package dingding

import (
	"github.com/blinkbean/dingtalk"
	"github.com/gin-gonic/gin"
	"promethues_alertmanager_webhook/common/code"
	"promethues_alertmanager_webhook/common/object"
	"promethues_alertmanager_webhook/config"
	"promethues_alertmanager_webhook/dto"
	"promethues_alertmanager_webhook/util"
)

func PostMarkdown(ctx *gin.Context) {
	var requestBody dto.AlertPayload
	err := ctx.ShouldBind(&requestBody)
	if err != nil {
		object.APIResponse(ctx, code.ErrValidation, err.Error())
		return
	}
	for _, alert := range requestBody.Alerts {
		dingTalkRequestBody, err := util.TextTemplateReader(alert, config.DingTalkTpl)
		if err != nil {
			object.APIResponse(ctx, code.InternalServerError, err.Error())
			return
		}
		dingTalkConfig := config.AppConf.Chat.DingTalk
		cli := dingtalk.InitDingTalkWithSecret(dingTalkConfig.AccessToken, dingTalkConfig.Secret)
		if err = cli.SendMarkDownMessage(alert.Annotations.Summary, dingTalkRequestBody); err != nil {
			object.APIResponse(ctx, code.InternalServerError, err.Error())
		}
	}

}
