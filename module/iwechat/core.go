package iwechat

import (
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
	mpOauth "github.com/chanxuehong/wechat.v2/mp/oauth2"
	"github.com/chanxuehong/wechat.v2/oauth2"
	"net/http"
)

var WeChatAccessTokenServer core.AccessTokenServer
var WeChatCallBackServer *core.Server
var WeChatClient *core.Client
var WeChatOauthClient *oauth2.Client
var WeChatJsTicketServer jssdk.TicketServer

type wechatConfig struct {
	AppId     string
	AppSecret string
}

var WechatConfig *wechatConfig

//初始化中控服务器
func InitAccessTokenServer(appid, secret string) {

	if WechatConfig == nil {
		WechatConfig = new(wechatConfig)
		WechatConfig.AppId = appid
		WechatConfig.AppSecret = secret
	}
	if WeChatAccessTokenServer == nil {
		WeChatAccessTokenServer = core.NewDefaultAccessTokenServer(appid, secret, nil)
	}
	t, err := WeChatAccessTokenServer.Token()
	if err != nil {
		logger.Error(err)
	}
	logger.Debug(t)
}

//初始化回调服务器
func InitCallBackServer(account, appid, token, aesKey string) {
	if WeChatCallBackServer == nil {
		initMux()
		WeChatCallBackServer = core.NewServer(account, appid, token, aesKey, mux, nil)
	}
}

//初始化客户端
func InitClient() {

	if WeChatClient == nil {
		WeChatClient = core.NewClient(WeChatAccessTokenServer, nil)
	}

}

func InitOAuthClient(appid, secret string) {
	if WeChatOauthClient == nil {
		WeChatOauthClient = new(oauth2.Client)
		WeChatOauthClient.Endpoint = mpOauth.NewEndpoint(appid, secret)
	}

}

func WeChatServerHandler(w http.ResponseWriter, r *http.Request) {
	WeChatCallBackServer.ServeHTTP(w, r, nil)
}

//依赖core.Client,最后启动
func InitJsTicketServer() {
	if WeChatJsTicketServer == nil {
		WeChatJsTicketServer = jssdk.NewDefaultTicketServer(WeChatClient)
	}
}
