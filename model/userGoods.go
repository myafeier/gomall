package model

import (
	"sync"
)

//用户级别与折扣的对应关系
type UserRankClassifyDiscount struct {
	Id        int64 `json:"id"`
	UserGroup int   `json:"user_group" xorm:"tinyint(2) index"` //用户组
	UserRank  int   `json:"user_rank" xorm:"tinyint(2) index"`  //用户级别
	Classify  int64 `json:"classify" xorm:"index"`              //分类id，设置某个分类时，其下所有的分类均为相同折扣
	Discount  int   `json:"discount" xorm:"tinyint(3)"`         // 折扣率 ?%
}

var cdSync = new(sync.RWMutex)
var CacheUserRankClassifyDiscount map[int]map[int]map[int64]int //用户级别与折扣的对应关系缓存,组，级别，分类

func InitUserRankClassifyDiscount() error {
	CacheUserRankClassifyDiscount = make(map[int]map[int]map[int64]int)

	t := new(UserRankClassifyDiscount)
	all, err := t.GetAll()
	if err != nil {
		logger.Error(err)
		return err
	}

	cdSync.RLock()
	defer cdSync.RUnlock()

	for _, v := range all {
		if _, ok := CacheUserRankClassifyDiscount[v.UserGroup]; !ok {
			CacheUserRankClassifyDiscount[v.UserGroup] = make(map[int]map[int64]int)
		}
		if _, ok := CacheUserRankClassifyDiscount[v.UserGroup][v.UserRank]; !ok {
			CacheUserRankClassifyDiscount[v.UserGroup][v.UserRank] = make(map[int64]int)
		}
		CacheUserRankClassifyDiscount[v.UserGroup][v.UserRank][v.Classify] = v.Discount
	}
	return nil
}

func (self *UserRankClassifyDiscount) GetAll() (result []*UserRankClassifyDiscount, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(UserRankClassifyDiscount)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}
	return

}
func (self *UserRankClassifyDiscount) GetOne() (bool, error) {
	return Db.Get(self)
}

//添加一个分类的时候需要对该分类下的所有子分类进行遍历
func (self *UserRankClassifyDiscount) Insert() (int64, error) {
	return Db.Insert(self)
}

//添加一个分类的时候需要对该分类下的所有子分类进行遍历
func (self *UserRankClassifyDiscount) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}
