package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"filter"
	"github.com/garyburd/redigo/redis"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
	"math/rand"
	"model"
	"module/iwechat"
	"strconv"
	"time"
)

const REDIS_SESSION_EXPIRE = 1 * 24 * 3600

//用户的微信信息表单结构
type WxInfo struct {
	AvatarUrl string `json:"avatarUrl,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Gender    int    `json:"gender,omitempty"`
	Language  string `json:"language,omitempty"`
	NickName  string `json:"nickName,omitempty"`
	Province  string `json:"province,omitempty"`
}
type WechatInfoForm struct {
	Code   string `json:"code"`
	WxInfo WxInfo `json:"userinfo"`
}

// /user/wechat/sign/up 微信用户注册
// code 换取 session_key
func UserWechatSignUpPostHandler(ctx *macaron.Context, sess session.Store, redisPool *redis.Pool) {

	requestBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		logger.Error(err)
		ctx.Status(400)
		return
	}
	request := new(WechatInfoForm)
	err = json.Unmarshal(requestBody, request)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}

	logger.Debug("Code:", request.Code)
	wxSession, err := iwechat.GetSession(request.Code)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}

	userSession := new(filter.Session)

	userSession.WechatSessionID = wxSession.SessionKey
	userSession.OpenID = wxSession.OpenID
	//生成sessionKEY
	rand.Seed(time.Now().Unix())
	logger.Debug("Login User OpenID:", userSession.OpenID)

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(userSession.WechatSessionID + strconv.Itoa(rand.Int())))
	cipher := md5Ctx.Sum(nil)
	userSession.SessionID = hex.EncodeToString(cipher)

	//检查此顾客是否已入库
	user := new(model.User)
	user.OpenId = wxSession.OpenID
	has, err := user.GetOne()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}

	user.WxAvatarPath = request.WxInfo.AvatarUrl
	user.WxCity = request.WxInfo.City
	user.WxCountry = request.WxInfo.Country
	user.WxNickName = request.WxInfo.NickName

	//如果已入库,且已是，直接登陆即可
	if has {
		_, err = user.Update()
		if err != nil {
			logger.Error(err)
			ctx.Status(500)
			return
		}

	} else { //如果未入库，入库，

		_, err = user.Insert()
		if err != nil {
			logger.Error(err)
			ctx.Status(500)
			return
		}
	}

	userSession.User = user
	userSessionJSON, err := json.Marshal(userSession)
	bytesBuffer := new(bytes.Buffer)
	bytesBuffer.Write(userSessionJSON)
	//写入redis
	redisConn := redisPool.Get()
	redisConn.Do("SET", filter.SESSION_PRE+userSession.SessionID, bytesBuffer.String())
	redisConn.Do("Expire", filter.SESSION_PRE+userSession.SessionID, REDIS_SESSION_EXPIRE)
	ctx.JSON(200, map[string]interface{}{"session": userSession.SessionID, "stat": user.Stat})
}

// {post} /user/wechat/sign/in 微信用户登录
// 此接口可以被filter/miniAppLogin过滤取代
func UserWechatSignInPostHandler(ctx *macaron.Context) {

}

// {post} /user/wechat/decrypt/data 微信用户信息的数据解密
func UserWechatDecryptDataPostHandler(ctx *macaron.Context) {

}

// {post} /user/sign/up 用户注册
func UserSignUpPostHandler(ctx *macaron.Context) {

}

//  {post} /user/sign/in 用户登录
func UserSignInPostHandler(ctx *macaron.Context) {

}

// {post} /user/sign/out 用户登出
func UserSignOutPostHandler(ctx *macaron.Context) {

}

// {post} /user/reset/password 修改密码
func UserPasswordResetPostHandler(ctx *macaron.Context) {

}

// {post} /user/info 保存用户信息
func UserInfoPostHandler(ctx *macaron.Context) {

}

// {get} /user/info 获取用户信息
func UserInfoGetHandler(ctx *macaron.Context) {

}

//获取自己的会员信息
//Get:/user/info/myself
func UserInfoOfSelf(ctx *macaron.Context) {

	userSession := GetUserSession(ctx)
	user := new(model.User)
	user.OpenId = userSession.OpenID
	has, err := user.GetOne()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	if has {

		ctx.JSON(200, user)
	} else {
		ctx.Status(204)
	}
}

//获取用户自己变动记录
func UserAccountLog(ctx *macaron.Context) {
	stat := ctx.ParamsInt("stat")
	page := ctx.ParamsInt("page")
	limit := ctx.ParamsInt("limit")

	logger.Debug("stat,page,limit:", stat, page, limit)
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}
	if stat < 1 {
		stat = 1
	}
	userSession := GetUserSession(ctx)

	m := new(model.UserAccountLogMP)
	m.Limit = limit
	m.Page = page
	err := m.GetListByStat(userSession.User.Id)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(200, m)

}

//更新用户的微信信息
func UserWechatInfoUpdate(ctx *macaron.Context, userInfoForm WechatInfoForm) {
	//logger.Debugf("%#v", userInfoForm)
	if userInfoForm.Code == "" || userInfoForm.WxInfo.NickName == "" {
		logger.Errorf("session wrong")
		ctx.Status(400)
		return
	}
	userSession := GetUserSession(ctx)
	u := new(model.User)
	u.OpenId = userSession.OpenID

	has, err := u.GetOne()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	if !has {
		logger.Errorf("%s", "No member found")
		ctx.Status(400)
		return
	}

	if userInfoForm.WxInfo.NickName != "" {
		u.WxNickName = userInfoForm.WxInfo.NickName
	}
	if userInfoForm.WxInfo.AvatarUrl != "" {
		u.WxAvatarPath = userInfoForm.WxInfo.AvatarUrl
	}
	if userInfoForm.WxInfo.City != "" {
		u.WxCity = userInfoForm.WxInfo.City
	}
	if userInfoForm.WxInfo.Province != "" {
		u.WxProvince = userInfoForm.WxInfo.Province
	}
	if userInfoForm.WxInfo.Country != "" {
		u.WxCountry = userInfoForm.WxInfo.Country
	}
	if userInfoForm.WxInfo.Gender != 0 {
		u.Sex = userInfoForm.WxInfo.Gender
	}
	_, err = u.Update()

	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}

}

//基本分页表单结构
type BasicLogByMPForm struct {
	//Stat int `form:"stat" binding:"Required"`      //订单状态
	PageNo  int64 `form:"page_no" binding:"OmitEmpty"` //第几页
	PerPage int64 `form:"per_page"`                    //每页多少
}
