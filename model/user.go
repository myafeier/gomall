package model

import (
	"errors"
	"github.com/go-xorm/xorm"
	"time"
)

const (
	RANK_1                = 1 //会员级别
	RANK_2                = 2
	RANK_3                = 3
	RANK_4                = 4
	RANK_5                = 5
	RANK_6                = 6
	RANK_1_COMMISION      = 0.05  //配送提成：5%
	RANK_PARENT_COMMISION = 0.001 //1/1000
	GROUP_CUSTOMER        = 1     // "终端用户组"
	GROUP_CATERERS        = 2     //"餐饮用户组"
	GROUP_EMPLOYEMPLOYEE  = 3     //"员工组"
)

type UserMP struct {
	Total   int64   `json:"total"`
	Pages   int     `json:"pages"`
	Page    int     `json:"page"`
	Limit   int     `json:"limit"`
	Content []*User `json:"content"`
}

func (self *UserMP) GetList(rank, stat int) (err error) {

	return
}

type User struct {
	Id              int64   `json:"id"`
	Guid            string  `json:"guid" xorm:"varchar(50) index"`
	Stat            int     `json:"stat" xorm:"tinyint(2) index"` //1正常，-1停用
	Mobile          string  `json:"mobile" xorm:"varchar(30) index"`
	Group           int     `json:"group" xorm:"tinyint(2) index"`            //用户组
	Rank            int64   `json:"rank,omitempty" xorm:"index"`              //会员级别，1级为普通消费者，1-6级为合作伙伴。
	TotalAmount     float32 `json:"total_amount" xorm:"decimal(10,2)"`        //总共赚的钱
	AvailableAmount float32 `json:"available_amount" xorm:"decimal(10,2)"`    //还没有提走的钱
	ZoneId          int64   `json:"zone_id,omitempty" xorm:"index"`           //隶属的区域会员ID
	RecommendUserId int64   `json:"recommend_user_id,omitempty" xorm:"index"` //上线ID-》谁推荐的
	StoreId         int64   `json:"store_id,omitempty" xorm:"index"`          //管理的配送店ID，仅有1级会员拥有
	OpenId          string  `json:"open_id,omitempty" xorm:"varchar(100) index"`
	UnionId         string  `json:"union_id,omitempty" xorm:"varchar(100)"`
	ParentId        int64   `json:"parent_id" xorm:"index"` //上级用户ID
	Sex             int     `json:"sex" xorm:"tinyint(2)"`  //用户性别
	TrueName        string  `json:"true_name" xorm:"varchar(30)"`
	WxNickName      string  `json:"nick_name,omitempty" xorm:"varchar(100)"` //微信昵称
	StatF           string  `json:"stat_f" xorm:"-"`
	WxCity          string  `json:"city,omitempty" xorm:"varchar(100)"`
	WxProvince      string  `json:"province,omitempty" xorm:"varchar(100)"`
	WxCountry       string  `json:"country,omitempty" xorm:"varchar(100)"`
	WxAvatarPath    string  `json:"avatar_path,omitempty" xorm:"varchar(200)"` //否	微信头像图片路径
	RankF           string  `json:"rank_f" xorm:"-"`
	GroupF          string  `json:"group_f" xorm:"-"`
}

type UserAccountLogMP struct {
	Total   int64             `json:"total"`
	Pages   int               `json:"pages"`
	Page    int               `json:"page"`
	Limit   int               `json:"limit"`
	Content []*UserAccountLog `json:"content"`
}

//会员资金账户变动日志
type UserAccountLog struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id" xorm:"index"`        //会员ID
	Amount     float32   `json:"amount" xorm:"decimal(10,2)"` //发生金额
	Before     float32   `json:"before" xorm:"decimal(10,2)"` //发生前
	After      float32   `json:"after" xorm:"decimal(10,2)"`  //发生后
	Action     string    `json:"action" xorm:"varchar(30)"`   //动作
	AtTime     time.Time `json:"at_time"`                     //发生时间
	ByUserID   int64     `json:"by_user_id"`                  //操作人
	AtTimeF    string    `json:"at_time_f" xorm:"-"`
	ByUserName string    `json:"by_user_name" xorm:"-"`
}

func (self *User) GetOne() (has bool, err error) {
	has, err = Db.Get(self)
	formatUserInfo(self)
	return
}

func (self *User) LockForUpdate(session *xorm.Session) (has bool, err error) {
	return session.ForUpdate().Get(self)
}
func (self *User) TUpdate(session *xorm.Session) (num int64, err error) {
	return session.ID(self.Id).Update(self)
}

func (self *UserAccountLog) TInsert(session *xorm.Session) (num int64, err error) {
	return session.Insert(self)
}

//递归发放派送酬金
func (self *User) TAddDispatchCommision(session *xorm.Session, action string, missionPay float32) (err error) {

	has, err := self.LockForUpdate(session)
	if err != nil {
		logger.Error(err)
		return
	}
	if !has {
		err = errors.New("miss user!")
		logger.Error(err)
		return
	}
	userLog := new(UserAccountLog)
	userLog.UserId = self.Id
	userLog.Action = action
	userLog.Before = self.AvailableAmount

	var addMoney float32
	if self.Rank == RANK_1 {
		addMoney = missionPay * RANK_1_COMMISION
	} else {
		addMoney = missionPay * RANK_PARENT_COMMISION
	}
	if addMoney < 0.01 {
		addMoney = 0.01
	}
	userLog.Amount = addMoney

	self.TotalAmount = self.TotalAmount + addMoney
	self.AvailableAmount = self.AvailableAmount + addMoney

	userLog.After = self.AvailableAmount
	num, err := self.TUpdate(session)
	if err != nil {
		logger.Error(err)
		return
	}
	if num != 1 {
		err = errors.New("update user error!")
		logger.Error(err)
		return
	}

	//增加日志
	userLog.AtTime = time.Now()
	_, err = userLog.TInsert(session)
	if err != nil {
		logger.Error(err)
		return
	}
	if self.ParentId != 0 {
		u := new(User)
		u.Id = self.ParentId
		err = self.TAddDispatchCommision(session, action, missionPay)
		if err != nil {
			logger.Error(err)
			return
		}
	}

	return
}

func formatUserInfo(u *User) {
	if u.Group == GROUP_CUSTOMER {
		switch u.Rank {
		case 1:
			u.RankF = "非会员"
		case 2:
			u.RankF = "铜牌会员"
		case 3:
			u.RankF = "银牌会员"
		case 4:
			u.RankF = "金牌会员"
		case 5:
			u.RankF = "钻石会员"
		case 6:
			u.RankF = "黑金会员"
		}
	} else if u.Group == GROUP_CATERERS {
		switch u.Rank {
		case 1:
			u.RankF = "待审核"
		case 2:
			u.RankF = "铜牌会员"
		case 3:
			u.RankF = "银牌会员"
		case 4:
			u.RankF = "金牌会员"
		}
	}

}
func (self *User) Insert() (num int64, err error) {
	return Db.Insert(self)
}

func (self *User) Update() (num int64, err error) {
	return Db.Id(self.Id).Update(self)
}

func (self *UserAccountLog) Count() (num int64, err error) {
	return Db.Count(self)
}

//分页获取任务列表
func (self *UserAccountLogMP) GetListByStat(userID int64) (err error) {
	ul := new(UserAccountLog)
	ul.UserId = userID
	self.Total, _ = ul.Count()

	logger.Debugf("%#v", ul)
	rows, err := Db.OrderBy("id desc").Limit(self.Limit, (self.Page-1)*self.Limit).Rows(ul)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		t := new(UserAccountLog)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}

		rebuildUserAccountLog(t)

		self.Content = append(self.Content, t)
	}

	tp := int(self.Total) / self.Limit
	if self.Total == 0 {
		self.Pages = 0
	} else if int(self.Total) <= self.Limit {
		self.Pages = 1
	} else if tp%self.Limit == 0 {
		self.Pages = tp
	} else {
		self.Pages = tp + 1
	}
	return
}

func rebuildUserAccountLog(ul *UserAccountLog) {
	if !ul.AtTime.IsZero() {
		ul.AtTimeF = ul.AtTime.Format("2006-01-02 15:04:05")
	}
}
