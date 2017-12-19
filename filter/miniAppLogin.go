package filter

import (
	"encoding/json"
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"gopkg.in/macaron.v1"
	"html/template"
	"model"
	"strings"
)

const SESSION_PRE = "huiwanjia_sessions_"

type Session struct {
	ID              string      `json:"id"`
	SessionID       string      `json:"session_id,omitempty"`        //顾客的本地会话ID
	WechatSessionID string      `json:"wechat_session_id,omitempty"` //顾客的微信会话ID
	OpenID          string      `json:"open_id"`                     //顾客的微信openID
	User            *model.User `json:"user,omitempty"`
}

//检查微信小程序的登陆状态及权限
//
func WxMiniAppLoginFilter(ctx *macaron.Context, redisPool *redigo.Pool) {
	authorization := ctx.Req.Header.Get("Authorization")
	tokens := strings.Split(authorization, " ")
	if len(tokens) < 2 || tokens[1] == "" {
		fmt.Println("Token is null,auth:", authorization)
		ctx.JSON(302, "Need Login")
		return
	}
	token := template.HTMLEscapeString(tokens[1])
	if token == "" {
		fmt.Println("Token is invalid:", tokens[1])
		ctx.JSON(302, "Need Login")
		return
	}
	redis := redisPool.Get()
	defer redis.Close()

	sessionJSON, err := redigo.Bytes(redis.Do("GET", SESSION_PRE+token))
	if err != nil {
		fmt.Println(err)
		if err == redigo.ErrNil {
			fmt.Println("No session In redis,Need login")
			ctx.JSON(302, "Need Login")
		} else {
			ctx.Status(500)
		}
		return
	}

	session := new(Session)
	err = json.Unmarshal(sessionJSON, session)
	if err != nil {
		fmt.Println(err)
		ctx.Status(500)
		return
	}
	//获取用户信息
	u := new(model.User)
	u.Id = session.User.Id
	has, err := u.GetOne()

	if err != nil {
		fmt.Println(err)
		ctx.Status(500)
		return
	}
	if !has {
		ctx.JSON(302, "Need Regist")
		return
	}
	//
	if u.Stat != 1 {
		ctx.JSON(401, "user account denied")
		return
	}
	//小P不允许终端用户登陆
	if u.Group == model.GROUP_CUSTOMER {
		ctx.JSON(401, "Group customer denied!")
		return
	}

	//权限判断,餐饮组级别为1的待审核会员不允许使用
	if u.Group == model.GROUP_CATERERS && u.Rank == model.RANK_1 {
		ctx.JSON(401, "Authorization required")
		return
	}

	ctx.Data["userInfo"] = u
	return

}
