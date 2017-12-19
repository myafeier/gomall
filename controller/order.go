package controller

import (
	"errors"
	"fmt"
	"gopkg.in/macaron.v1"
	"model"
	"time"
)

//分页获取我的订单配送记录表单
type OrderDispatchListByMPOfSelfJSON struct {
	Stat  int   `form:"stat" binding:"Required"`  //订单状态
	Page  int64 `form:"page" binding:"OmitEmpty"` //第几页
	Limit int64 `form:"limit"`                    //每页多少
}

//分页获取我的订单配送记录
//分状态获取
// stat=1,待配送；stat=2,已配送；stat=3，已取消
func OrderDispatchListByMPOfSelf(ctx *macaron.Context) {

	stat := ctx.ParamsInt("stat")
	page := ctx.ParamsInt("page")
	limit := ctx.ParamsInt("limit")

	fmt.Println("stat,page,limit:", stat, page, limit)
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

	m := new(model.OrderDispatchMissionByMP)
	m.Limit = limit
	m.Page = page
	err := m.GetListByStat(userSession.User.Id, stat)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(200, m)
}

//配送逻辑
func OrderMissionDispatch(ctx *macaron.Context) {
	missionID := ctx.ParamsInt64("id")
	if missionID < 1 {
		ctx.JSON(400, "Invalid mission id!")
		return
	}
	userSession := GetUserSession(ctx)
	if userSession == nil {
		ctx.JSON(401, "Need Login!")
		return
	}

	missionTemp := new(model.OrderDispatchMission)
	missionTemp.Id = missionID
	missionTemp.Stat = 1
	missionTemp.DispatchUserId = userSession.User.Id
	has, err := missionTemp.GetOne()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	if !has {
		ctx.JSON(400, "无对应的派送单")
		return
	}

	//获取订单的TotalMoney
	order := new(model.Order)
	order.ID = missionTemp.OrderId
	has, err = order.GetOne()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	if !has {
		ctx.JSON(400, "无对应的订单")
		return
	}

	t := model.Db.NewSession()
	err = t.Begin()
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	mission := new(model.OrderDispatchMission)
	mission.Id = missionID
	mission.Stat = 1
	mission.DispatchUserId = userSession.User.Id
	has, err = mission.LockForUpdate(t)
	if err != nil {
		logger.Error(err)
		ctx.Status(500)
		return
	}
	if !has {
		t.Rollback()
		ctx.JSON(400, "无对应的派送单")
		return
	}

	mission.Stat = 2
	mission.DispatchTime = time.Now()
	num, err := mission.TUpdate(t)
	if err != nil {
		logger.Error(err)
		t.Rollback()
		ctx.Status(500)
		return
	}
	if num != 1 {
		err = errors.New("no mission had be updated")
		logger.Error(err)
		t.Rollback()
		ctx.Status(500)
		return
	}

	//发放配送酬金
	user := new(model.User)
	user.Id = mission.DispatchUserId
	err = user.TAddDispatchCommision(t, "配送酬金", order.TotalMoney)
	if err != nil {
		logger.Error(err)
		t.Rollback()
		ctx.Status(500)
		return
	}

	err = t.Commit()

	if err != nil {
		logger.Error(err)
		t.Rollback()
		ctx.Status(500)
		return
	}
	ctx.Status(200)
}
