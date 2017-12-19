package model

import "time"

var MemberPointRule=map[string]int64{
	"newRegister":200,         //注册会员就送200
	"recommendedMember0":10,	//推荐的0级会员送10
	"recommendedMember1":5000,	//推荐的1级会员
	"recommendedMember2":10000,	//推荐的2级会员
	"recommendedMember3":30000,	//推荐的3级会员
	"recommendedMember4":100000,	//推荐的4级会员
	"recommendedMember5":200000,	//推荐的5级会员
	"recommendedMember6":400000,	//推荐的6级会员
}

//会员积分，废弃
type MemberPoint struct {
	ID int64
	MemberID int64 //会员ID
	TotalValue	int64 //是	历史积分
	AvailableValue	int64 //是	可用积分
}

type MemberPointLogMP struct {
	Total   int64                 `json:"total"`
	Pages   int                   `json:"pages"`
	Page    int                   `json:"page"`
	Content []*MemberPointLog `json:"content"`
}
//会员资金账户变动日志
type MemberPointLog struct {
	ID int64 `json:"id"`
	MemberID int64 `json:"member_id"`//会员ID
	MemberPointID int64 `json:"member_point_id"`//积分账户ID
	Amount int64 `json:"amount"`//发生数量
	Before int64 `json:"before"`//发生前
	After int64 `json:"after"`//发生后
	Action int `json:"action"`//动作
	AtTime time.Time `json:"at_time"`//发生时间
	ByMemberID int64 `json:"by_member_id"`//操作人
	ActionF string `json:"action_f"`//
	AtTimeF string `json:"at_time_f"`
	ByMemberName string `json:"by_member_name"`
}

