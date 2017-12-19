package iwechat

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
)

const UniformReplyMessage = "您好!欢迎访问云南省医疗美容文饰分会服务号，若需要即时回复,请您致电0871-65710050。"

func messageHandler(ctx *core.Context) {

	//if ctx.MixedMsg.Content == "test" {
	//	msg := response.NewText(ctx.MixedMsg.FromUserName, ctx.MixedMsg.ToUserName, time.Now().Unix(), oauth2.AuthCodeURL(config.WeChatConfig.WxAppid, config.SiteCommonConfig.SiteUrl+"/goodstest/", "snsapi_userinfo", "state"))
	//	ctx.ResponseWriter.Header().Set("Content-Type", "text/xml")
	//	ctx.RawResponse(msg)
	//
	//} else {
	//	msg := response.NewText(ctx.MixedMsg.FromUserName, ctx.MixedMsg.ToUserName, time.Now().Unix(), UniformReplyMessage)
	//	ctx.ResponseWriter.Header().Set("Content-Type", "text/xml")
	//	ctx.RawResponse(msg)
	//}

}

//消息处理接口
type ImessageHandler interface {
	SendTemplateMessage(openId string, templateID int, message map[string]string) error
}
