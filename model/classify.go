package model

import (
	"sync"
)

const (
	MAX_LEVEL = 3
)

// 商品分类
type Classify struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" xorm:"varchar(100)"`
	Level    int    `json:"level" xorm:"tinyint(1)"`  //层级
	Icon     string `json:"icon" xorm:"varchar(200)"` //分类图标
	ParentId int64  `json:"parent_id" xorm:"index"`
}

//带层级结构的分类
type Classifys struct {
	Level int          `json:"level"`
	Self  *Classify    `json:"self"`  //当前级
	Child []*Classifys `json:"child"` //子级
}

//读写安全的分类结构
type SafeClassify struct {
	sync.RWMutex
	Map         map[int64]*Classify
	LevelMap    map[int][]*Classify
	LevelStruct []*Classifys //结构化的分类
}

var CacheClassify *SafeClassify

//初始化缓存
func InitClassify() error {

	CacheClassify = new(SafeClassify)
	CacheClassify.LevelMap = make(map[int][]*Classify)
	CacheClassify.Map = make(map[int64]*Classify)
	s := new(Classify)
	result, err := s.GetAll()

	if err != nil {
		logger.Error(err)
		return err
	}
	for _, v := range result {
		CacheClassify.Map[v.Id] = v
		CacheClassify.LevelMap[v.Level] = append(CacheClassify.LevelMap[v.Level], v)
	}
	//开始结构化存储

	for i := 1; i <= MAX_LEVEL; i++ {

		value, ok := CacheClassify.LevelMap[i]
		if !ok {
			continue
		}
		// logger.Debug("level:", i)
		for _, v := range value {
			t := new(Classifys)
			t.Level = v.Level
			t.Self = v
			if i == 1 {
				CacheClassify.LevelStruct = append(CacheClassify.LevelStruct, t)
				// logger.Debug("level:", i, "now:", CacheClassify.LevelStruct)
			} else {
				recursionInsert(CacheClassify.LevelStruct, t)
				// logger.Debug("re level:", i, "now:", CacheClassify.LevelStruct)
			}
		}

	}
	return nil
}

//对树状结构分类做递归插入
func recursionInsert(last []*Classifys, target *Classifys) {
	for k, v := range last {

		if v.Self.Id == target.Self.ParentId {
			last[k].Child = append(last[k].Child, target)
			return
		}

		if v.Child != nil {
			recursionInsert(last[k].Child, target)
		}
	}
}

//获取某个分类的所有子类切片
func FindAllClassifyChildIdSlice(searchId int64) (result []int64) {
	return recursionFindChildSlice(CacheClassify.LevelStruct, searchId, false)
}

//递归获取一个分类的所有子分类
func recursionFindChildSlice(target []*Classifys, searchId int64, start bool) (result []int64) {

	for _, v := range target {
		if start == false { //还没有找到起点

			if searchId == v.Self.Id { //发现起点，开始插入逻辑
				result = append(result, searchId)
				start = true
				if v.Child != nil {

					t := recursionFindChildSlice(v.Child, searchId, start)
					result = append(result, t...)
				}
				break //在某一轮中找到了起点钻入遍历即可。就不需要再同级便利了
			} else { //继续在子节点中寻找起点
				if v.Child != nil {
					t := recursionFindChildSlice(v.Child, searchId, start)
					result = append(result, t...)
				}
			}
		} else { //已经找到起点，便利即可

			result = append(result, v.Self.Id)
			if v.Child != nil {
				t := recursionFindChildSlice(v.Child, searchId, start)
				result = append(result, t...)
			}
		}
	}

	return result
}

func (self *Classify) GetAll() (result []*Classify, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(Classify)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return

}
func (self *Classify) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *Classify) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *Classify) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

//删除分类时需要将分类折扣对应删除。
func (self *Classify) Delete() (int64, error) {
	return Db.Id(self.Id).Delete(self)
}
