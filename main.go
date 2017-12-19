package main

import "gopkg.in/macaron.v1"
import (
	"config"
	"controller"
	"filter"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/go-macaron/binding"
	"github.com/go-macaron/session"
	"model"
	"module/iwechat"
	"os"
	"time"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config.Init(pwd + "/etc/config.json")

	//微信中控服务器启动
	iwechat.InitAccessTokenServer(config.SiteConfig.WX.WxAppID, config.SiteConfig.WX.WxSecret)
	//微信回调服务器启动
	iwechat.InitCallBackServer(config.SiteConfig.WX.WxAccount, config.SiteConfig.WX.WxAppID, config.SiteConfig.WX.WxToken, config.SiteConfig.WX.WxAesKey)
	////微信客户端初始化
	iwechat.InitClient()
	////微信Oauth客户端初始化
	iwechat.InitOAuthClient(config.SiteConfig.WX.WxAppID, config.SiteConfig.WX.WxSecret)
	////微信JSTicket Server初始化
	iwechat.InitJsTicketServer()

	model.InitDb("mysql", config.SiteConfig.MYSQL.MYSQL_HOST, config.SiteConfig.MYSQL.MYSQL_PORT, config.SiteConfig.MYSQL.MYSQL_USER, config.SiteConfig.MYSQL.MYSQL_PWD, config.SiteConfig.MYSQL.MYSQL_DB)

	//redis初始化
	//redisDbOption := redigo.DialDatabase(common.RedisDb)
	//redisPasswordOption:=redis.DialPassword(config.RedisPassword)

	//启动有赞TokenServer
	// youzan.InitTokenServer(config.SiteConfig.YOUZAN.Youzan_Client_Id, config.SiteConfig.YOUZAN.Youzan_Client_Secret, "", config.SiteConfig.YOUZAN.Youzan_Kdt_Id)

}
func main() {

	redisPool := &redigo.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redigo.Conn, error) { return redigo.Dial("tcp", config.SiteConfig.Redis) },
	}
	defer redisPool.Close()

	m := macaron.New()
	m.Map(redisPool)
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(session.Sessioner())
	m.Use(macaron.Static("public"))
	m.Use(macaron.Renderer())
	//路由定义
	{
		//api 相关
		m.Post("/api/address/default/:id", controller.AddressDefaultPostHandler)
		m.Get("/api/address/default", controller.AddressDefaultGetHandler)
		m.Get("/api/address", controller.AddressListGetHandler)
		m.Get("/api/address/:id", controller.AddressOneGetHandler)
		m.Post("/api/address", controller.AddressPostHandler)
		m.Put("/api/address/:id", controller.AddressPutHandler)
		m.Delete("/api/address/:id", controller.AddressDeleteHandler)

		m.Post("/api/user/wechat/sign/up", controller.UserSignUpPostHandler)
		m.Post("/api/user/wechat/sign/in", controller.UserSignInPostHandler)
		m.Post("/api/user/wechat/decrypt/data", controller.UserWechatDecryptDataPostHandler)
		m.Post("/api/user/sign/up", controller.UserSignUpPostHandler)
		m.Post("/api/user/sign/in", controller.UserSignInPostHandler)
		m.Post("/api/user/sign/out", controller.UserSignOutPostHandler)
		m.Post("/api/user/reset/password", controller.UserPasswordResetPostHandler)
		m.Post("/api/user/info", controller.UserInfoPostHandler)
		m.Get("/api/user/info", controller.UserInfoGetHandler)

		m.Get("/api/banner/:is_show", controller.BannerListAllGetHandler)
		m.Get("/api/banner/:id")
		m.Post("/api/banner")
		m.Put("/api/banner/:id")
		m.Delete("/api/banner/:id")

		m.Get("/api/classify/?:level", controller.ClassifyGetListHandler)
		m.Get("/api/classify/:id")
		m.Post("/api/classify")
		m.Put("/api/classify/:id")
		m.Delete("/api/classify/:id")
		m.Get("/api/goods", controller.GoodsListByPagination)

		//用户会员相关
		m.Any("/app/handler", iwechat.WeChatServerHandler)
		m.Get("/user/myself/info", filter.WxMiniAppLoginFilter, controller.UserInfoOfSelf)
		m.Post("/user/wechat/info/update", filter.WxMiniAppLoginFilter, binding.Bind(controller.WechatInfoForm{}), controller.UserWechatInfoUpdate)

		m.Get("/user/myself/account/log/:page/:limit", filter.WxMiniAppLoginFilter, controller.UserAccountLog)
		//订单相关
		m.Get("/user/order/dispatch/list/:stat/:page/:limit", filter.WxMiniAppLoginFilter, controller.OrderDispatchListByMPOfSelf)
		m.Get("/user/order/dispatch/complete/:id", filter.WxMiniAppLoginFilter, controller.OrderMissionDispatch)
		//后台管理
		m.Get("/admin/user/list/:stat/:rank/:page/:limit", controller.AdminUserList)
		m.Get("/admin/user/detail", controller.AdminUserDetail)
		m.Post("/admin/user/update", controller.AdminUserUpdate)
		m.Post("/admin/user/search", controller.AdminUserSearch)
		m.Get("/admin/user/withdraw/list", controller.AdminUserWithdrawsList)
		m.Post("/admin/user/withdraws/deal", controller.AdminUserWithdrawsDeal)

		m.Get("/admin/order_dispatch_mission/list/:stat/:userId/:page/:limit", controller.AdminOrderDispatchMissionList)

	}
	m.Run(9201)
}
