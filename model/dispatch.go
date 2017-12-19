package model

import (
	"github.com/go-xorm/xorm"
	"time"
)

type OrderDispatchMissionByMP struct {
	Total   int64                   `json:"total"`
	Pages   int                     `json:"pages"`
	Page    int                     `json:"page"`
	Limit   int                     `json:"limit"`
	Content []*OrderDispatchMission `json:"content"`
}

//订单配送任务
type OrderDispatchMission struct {
	Id              int64     `json:"id"`
	OrderId         int64     `json:"order_id" xorm:"index"`          //订单ID
	DispatchUserId  int64     `json:"dispatch_user_id" xorm:"index"`  //配送人的ID
	OrderedMemberId int64     `json:"ordered_member_id" xorm:"index"` //下订单人的ID
	Stat            int       `json:"stat" xorm:"tinyint(2) index"`   //状态,1:待配送，2，已配送，3，取消
	DispatchTime    time.Time `json:"dispatch_time" xorm:""`          //配送时间
	CreateTime      time.Time `json:"create_time" xorm:""`            //配送时间
	OrderInfo       *Order    `json:"order_info" xorm:"-"`
	CreateTimeF     string    `json:"create_time_f" xorm:"-"`
	DispatchTimeF   string    `json:"dispatch_time_f" xorm:"-"` //配送时间
	StatF           string    `json:"stat_f" xorm:"-"`
}

//分页获取任务列表
func (self *OrderDispatchMissionByMP) GetListByStat(userID int64, stat int) (err error) {
	om := new(OrderDispatchMission)
	om.Stat = stat
	om.DispatchUserId = userID
	self.Total, _ = om.Count()

	logger.Debugf("%#v", om)
	rows, err := Db.OrderBy("id desc").Limit(self.Limit, (self.Page-1)*self.Limit).Rows(om)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		t := new(OrderDispatchMission)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		tt := new(Order)
		tt.ID = t.OrderId
		has, err1 := tt.GetOneWithItems()
		if err1 != nil {
			logger.Error(err1)
			return
		}

		if has {
			t.OrderInfo = tt
		}
		err = rebuildOrderDispatchMission(t)
		if err != nil {
			logger.Error(err)
			return
		}
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

func rebuildOrderDispatchMission(t *OrderDispatchMission) (err error) {
	if t.Stat != 0 {
		switch t.Stat {
		case 1:
			t.StatF = "待配送"
		case 2:
			t.StatF = "已配送"
		case 3:
			t.StatF = "已取消"
		}
	}
	if !t.DispatchTime.IsZero() {
		t.DispatchTimeF = t.DispatchTime.Format("2006-01-02 15:04:05")
	}
	if !t.CreateTime.IsZero() {
		t.CreateTimeF = t.CreateTime.Format("2006-01-02 15:04:05")
	}
	return
}

func (self *OrderDispatchMission) Count() (total int64, err error) {
	return Db.Count(self)
}

func (self *OrderDispatchMission) GetOne() (has bool, err error) {
	return Db.Get(self)
}

func (self *OrderDispatchMission) Update() (num int64, err error) {
	return Db.Id(self.Id).Update(self)
}
func (self *OrderDispatchMission) Insert() (num int64, err error) {
	return Db.Insert(self)
}

func (self *OrderDispatchMission) LockForUpdate(session *xorm.Session) (has bool, err error) {
	return session.ForUpdate().Get(self)
}
func (self *OrderDispatchMission) TUpdate(session *xorm.Session) (num int64, err error) {
	return session.ID(self.Id).Update(self)
}
