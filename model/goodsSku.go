package model

import (
	"encoding/json"
	"sync"
	"time"
)

//商品库存单元表
type GoodsSku struct {
	Id                 int64               `json:"id"`
	Gid                int64               `json:"gid" xorm:"index"`                         //商品ID
	SkuUniqueCode      string              `json:"sku_unique_code" xorm:"varchar(100)"`      //库存唯一编码
	PropertiesNameJson string              `json:"properties_name_json" xorm:"varchar(255)"` //规格属性json串[{'k':1,'v':2},{'k':1,'v':2}],key: spec attr; value:spec value
	Price              float32             `json:"price" xorm:"decimal(10,2)"`               //价格, 单位元,精确到小数点后两位
	DiscountRate       int                 `json:"discount_rate" xorm:"tinyint(3)"`          //价格, 单位元,精确到小数点后两位
	Quantity           int                 `json:"quantity" xorm:""`                         //库存数量
	Created            time.Time           `json:"created" xorm:""`                          //创建时间
	Modified           time.Time           `json:"modified" xorm:""`                         //  修改时间
	SoldQuantity       int64               `json:"sold_quantity" xorm:"not null default 0"`  //销售数量
	PropertiesF        []map[string]string `json:"properties,omitempty" xorm:"-"`            //规格属性格式化结果
}

type SafeGoodsSku struct {
	sync.RWMutex
	ValueMap map[int64]*GoodsSku //key : GoodsSku id
}

var CacheSafeGoodsSku *SafeGoodsSku

func InitSafeGoodsSku() (err error) {
	s := new(GoodsSku)
	result, err := s.GetAll()
	if err != nil {
		logger.Error(err)
		return
	}
	CacheSafeGoodsSku = new(SafeGoodsSku)
	CacheSafeGoodsSku.ValueMap = make(map[int64]*GoodsSku)
	CacheSafeGoodsSku.RLock()
	defer CacheSafeGoodsSku.RUnlock()
	for _, v := range result {
		CacheSafeGoodsSku.ValueMap[v.Id] = v
	}
	return
}

//商品规格属性对应值
type GoodsSkuSpecMapping struct {
	Id   int64     `json:"id"`
	GId  int64     `json:"g_id" xorm:"index"`  //商品ID
	AId  int64     `json:"a_id"  xorm:"index"` //规格名ID
	VId  int64     `json:"v_id"  xorm:"index"` //规格值id
	SId  int64     `json:"s_id"  xorm:"index"` //sku id
	AIdF string    `json:"a_id_f"  xorm:"-"`
	VIdF string    `json:"v_id_f"  xorm:"-"`
	SIdF *GoodsSku `json:"s_id_f" xorm:"-"` //sku
}

//格式化规格JSON
func formatGoodsSkuProperties(t *GoodsSku) (err error) {
	if t.PropertiesNameJson != "" {
		p := make([]map[int64]int64, 10)
		err = json.Unmarshal([]byte(t.PropertiesNameJson), p)
		if err != nil {
			return
		}
		t.PropertiesF = make([]map[string]string, 10)
		for _, v := range p {
			for kk, vv := range v {
				var kkString, vvString string
				kkStruct, ok := CacheSafeGoodsSpecAttr.ValueMap[kk]
				if ok {
					kkString = kkStruct.Name
				}
				vvStruct, ok := CacheSafeGoodsSpecValue.ValueMap[vv]
				if ok {
					vvString = vvStruct.Value
				}
				t.PropertiesF = append(t.PropertiesF, map[string]string{kkString: vvString})
			}

		}
	}

	return
}

func formatGoodsSkuSpecMapping(t *GoodsSkuSpecMapping) {
	if t.AId != 0 {
		s, ok := CacheSafeGoodsSpecAttr.ValueMap[t.AId]
		if ok {
			t.AIdF = s.Name
		}
	}
	if t.VId != 0 {
		s, ok := CacheSafeGoodsSpecValue.ValueMap[t.VId]
		if ok {
			t.VIdF = s.Value
		}
	}
	if t.SId != 0 {
		t.SIdF, _ = CacheSafeGoodsSku.ValueMap[t.SId]
	}
}

func (self *GoodsSku) GetAllWithFormatedSpec() (result []*GoodsSku, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSku)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		err = formatGoodsSkuProperties(t)
		if err != nil {
			logger.Error(err)
			return
		}

		result = append(result, t)
	}
	return

}

func (self *GoodsSku) GetAll() (result []*GoodsSku, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSku)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return

}
func (self *GoodsSku) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *GoodsSku) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *GoodsSku) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

func (self *GoodsSku) Delete() (int64, error) {
	return Db.Id(self.Id).Delete(self)
}

func (self *GoodsSkuSpecMapping) GetAllFormated() (result []*GoodsSkuSpecMapping, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSkuSpecMapping)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		formatGoodsSkuSpecMapping(t)
		result = append(result, t)
	}
	return

}

func (self *GoodsSkuSpecMapping) GetAll() (result []*GoodsSkuSpecMapping, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsSkuSpecMapping)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return

}
func (self *GoodsSkuSpecMapping) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *GoodsSkuSpecMapping) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *GoodsSkuSpecMapping) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

func (self *GoodsSkuSpecMapping) Delete() (int64, error) {
	return Db.Id(self.Id).Delete(self)
}
