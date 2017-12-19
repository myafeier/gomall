package model

import (
	"sync"
)

// 商品规格属性
type GoodsSpecAttr struct {
	Id   int64  `json:"id" `
	Name string `json:"name" xorm:"varchar(255)"` //规格属性名称
	Desc string `json:"desc" xorm:"varchar(255)"` //规格属性描述
}

//常用商品规格属性
type GoodsSpecValue struct {
	Id    int64  `json:"id"`
	AId   int64  `json:"a_id" xorm:"index"`         //规格属性ID
	Value string `json:"value" xorm:"varchar(255)"` //规格值
	Desc  string `json:"desc" xorm:"varchar(255)"`  //规格值描述
	Logo  string `json:"logo" xorm:"varchar(255)"`  //规格属性图片
}

// 商品标签对应
// 暂时不做
type GoodsTagMapping struct {
}

type SafeGoodsSpecAttr struct {
	sync.RWMutex
	ValueMap map[int64]*GoodsSpecAttr
}
type SafeGoodsSpecValue struct {
	sync.RWMutex
	ValueMap map[int64]*GoodsSpecValue
	KVMap    map[int64][]*GoodsSpecValue //key AID，属性ID
}

var CacheSafeGoodsSpecAttr *SafeGoodsSpecAttr
var CacheSafeGoodsSpecValue *SafeGoodsSpecValue

func InitCacheSafeGoodsSpecAttr() error {
	t := new(GoodsSpecAttr)
	result, err := t.GetAll()
	if err != nil {
		logger.Error(err)
		return err
	}

	CacheSafeGoodsSpecAttr = new(SafeGoodsSpecAttr)
	CacheSafeGoodsSpecAttr.ValueMap = make(map[int64]*GoodsSpecAttr)
	CacheSafeGoodsSpecAttr.RLock()
	defer CacheSafeGoodsSpecAttr.RUnlock()
	for _, v := range result {
		CacheSafeGoodsSpecAttr.ValueMap[v.Id] = v
	}
	return nil

}
func InitCacheSafeGoodsSpecValue() error {
	t := new(GoodsSpecValue)
	result, err := t.GetAll()
	if err != nil {
		logger.Error(err)
		return err
	}

	CacheSafeGoodsSpecValue = new(SafeGoodsSpecValue)
	CacheSafeGoodsSpecValue.KVMap = make(map[int64][]*GoodsSpecValue)
	CacheSafeGoodsSpecValue.ValueMap = make(map[int64]*GoodsSpecValue)

	CacheSafeGoodsSpecValue.RLock()
	defer CacheSafeGoodsSpecValue.RUnlock()
	for _, v := range result {
		CacheSafeGoodsSpecValue.ValueMap[v.Id] = v
		CacheSafeGoodsSpecValue.KVMap[v.AId] = append(CacheSafeGoodsSpecValue.KVMap[v.AId], v)
	}
	return nil

}

func (self *GoodsSpecAttr) GetAll() (result []*GoodsSpecAttr, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSpecAttr)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}
	return

}
func (self *GoodsSpecAttr) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *GoodsSpecAttr) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *GoodsSpecAttr) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

func (self *GoodsSpecValue) GetAll() (result []*GoodsSpecValue, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSpecValue)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}
	return

}
func (self *GoodsSpecValue) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *GoodsSpecValue) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *GoodsSpecValue) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}
