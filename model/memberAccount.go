package model

import "time"


//会员资金账户,废弃
type MemberAccount struct {
	ID int64
	MemberID int64 //会员ID
	TotalValue	int64 //是	历史储值
	AvailableValue	int64 //是	可用储值
}

type MemberAccountLogMP struct {
	Total   int64                 `json:"total"`
	Pages   int                   `json:"pages"`
	Page    int                   `json:"page"`
	Limit   int `json:"limit"`
	Content []*MemberAccountLog 	`json:"content"`
}

//会员资金账户变动日志
type MemberAccountLog struct {
	ID int64 `json:"id"`
	MemberID int64 `json:"member_id"`//会员ID 
	MemberAccountID int64 `json:"member_account_id"`//资金账户ID
	Amount int64 `json:"amount"`//发生金额
	Before int64 `json:"before"`//发生前
	After int64 `json:"after"`//发生后
	Action int `json:"action"`//动作
	AtTime time.Time `json:"at_time"`//发生时间
	ByMemberID int64 `json:"by_member_id"`//操作人
	ActionF string `json:"action_f"`//
	AtTimeF string `json:"at_time_f"`
	ByMemberName string `json:"by_member_name"`
}