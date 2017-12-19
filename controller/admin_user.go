package controller

import (
	"gopkg.in/macaron.v1"
	"model"
)

//查询所有的用户列表
func AdminUserList(ctx *macaron.Context) {

	//stat/:rank/:page/:limit
	stat := ctx.ParamsInt("stat")
	rank := ctx.ParamsInt("rank")
	page := ctx.ParamsInt("page")
	limit := ctx.ParamsInt("limit")
	logger.Debug("stat|rank|page|limit:", stat, "|", rank, "|", page, "|", limit)
	m := new(model.UserMP)
	ctx.JSON(200, m)

}

//修改用户信息
func AdminUserUpdate(ctx *macaron.Context) {

}

//添加用户
func AdminUserAdd(ctx *macaron.Context) {

}

//查询用户详情
func AdminUserDetail(ctx *macaron.Context) {

}

//搜索用户
func AdminUserSearch(ctx *macaron.Context) {

}

//查询用户提醒申请列表
func AdminUserWithdrawsList(ctx *macaron.Context) {

}

//暂时不做
//审批用户提现申请
func AdminUserWithdrawsDeal(ctx *macaron.Context) {

}

//暂时不做
//删除用户
func AdminUserDelete(ctx *macaron.Context) {
}
