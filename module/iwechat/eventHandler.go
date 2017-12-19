package iwechat

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
)

const SCANEvent = core.EventType("SCAN")
const SUBSCRIBEEvent = core.EventType("subscribe")

//事件处理器
func eventHandler(ctx *core.Context) {
	if ctx.MixedMsg.EventType == SCANEvent || ctx.MixedMsg.EventType == SUBSCRIBEEvent { //扫码事件
		//c := new(model.Customer)
		//c.OpenId = ctx.MixedMsg.FromUserName
		//has, err := c.GetOneSimple()
		//if err != nil {
		//	logger.Error(err)
		//	ctx.ResponseWriter.WriteHeader(500)
		//}
		//
		//if !has { //如果系统中没有记录，或者普通顾客，根据扫码 加入系统用户
		//
		//	if has {
		//		_, err = c.Update()
		//	} else {
		//		_, err = c.Insert()
		//	}
		//
		//	if err != nil {
		//		logger.Error(err)
		//		ctx.ResponseWriter.WriteHeader(500)
		//	}
		//
		//	csl := new(model.CustomerSubscribeLog)
		//	csl.Uid = c.Id
		//	if ctx.MixedMsg.EventKey == "g1" || ctx.MixedMsg.EventKey == "qrscene_g1" {
		//		csl.FromChannel = "g1"
		//	} else if strings.Contains(ctx.MixedMsg.EventKey, "p") {
		//		csl.FromChannel = "bookmark"
		//	}
		//	csl.AddTime = time.Now()
		//	_, err = csl.Insert()
		//	if err != nil {
		//		logger.Error(err)
		//		ctx.ResponseWriter.WriteHeader(500)
		//	}
		//}

		//生成微信菜单

		//顾客来自地铁广告牌
		//if ctx.MixedMsg.EventKey == "g1" || ctx.MixedMsg.EventKey == "qrscene_g1" {
		//	article1 := response.Article{Title: "《眉语匠心》发售了", Description: "详情请查看", PicURL: config.SiteCommonConfig.SiteUrl + "/images/book_thumb.jpg", URL: oauth2.AuthCodeURL(config.WeChatConfig.WxAppid, config.SiteCommonConfig.SiteUrl+"/goods/", "snsapi_userinfo", "state")}
		//	news := response.NewNews(ctx.MixedMsg.FromUserName, ctx.MixedMsg.ToUserName, time.Now().Unix(), []response.Article{article1})
		//	result, err := xml.Marshal(news)
		//	if err != nil {
		//		logger.Error(err)
		//		ctx.ResponseWriter.WriteHeader(500)
		//	}
		//	ctx.ResponseWriter.Header().Set("Content-Type", "text/xml")
		//	ctx.ResponseWriter.Write(result)
		//
		//} else if ctx.MixedMsg.EventKey == "external" || ctx.MixedMsg.EventKey == "qrscene_external" {
		//
		//}

		return
	}

}
