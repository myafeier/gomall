package model

import (
	"fmt"
	"sync"
	"time"
)

//商品状态
var GoodsStatMap = map[int]string{
	1:  "上架中",
	-1: "已下架",
	-2: "已删除",
}

type Goods struct {
	Id             int64                  `json:"id"`
	Name           string                 `json:"name" xorm:"varchar(100) notnull index"`                               //商品名称
	PinYin         string                 `json:"-" xorm:"varchar(255) not null"`                                       //拼音
	Stat           int                    `json:"stat,omitempty" xorm:"tinyint(2) notnull index"`                       //商品状态
	Classify       int64                  `json:"classify,omitempty" xorm:"notnull index"`                              //商品分类
	ChannelId      int                    `json:"channel_id,omitempty" xorm:"tinyint(2) notnull index"`                 //来源渠道ID                 //商品销售渠道 GoodsChannel
	Price          float32                `json:"price,omitempty"  xorm:"decimal(10,2)"`                                //价格范围,价格的最终确定是由sku确定                //零售价
	DiscountRate   int                    `json:"discount_rate,omitempty" xorm:"tinyint(3) notnull "`                   //默认零售折扣
	ExpireDay      time.Time              `json:"expire_day,omitempty"  xorm:""`                                        //购买后过期时间,单位为天
	Point          int                    `json:"point,omitempty"  xorm:"notnull default 0"`                            //积分
	SoldQuantity   int64                  `json:"sold_quantity" xorm:"not null default 0"`                              //销售数量
	Description    string                 `json:"description,omitempty" xorm:"text notnull "`                           //描述
	Logo           string                 `json:"logo,omitempty" xorm:"varchar(100) notnull default  ''"`               //图标
	CurrentVersion int                    `json:"current_version,omitempty" xorm:"smallint(4) notnull default 0 index"` //当前版本
	PriceF         string                 `json:"price_f,omitempty" xorm:"-"`
	DiscountRateF  string                 `json:"discount_rate_f,omitempty" xorm:"-"`
	StatF          string                 `json:"stat_f,omitempty" xorm:"-"`
	ChannelIdF     string                 `json:"channel_id_f,omitempty" xorm:"-"`
	Images         []*GoodsImages         `json:"images,omitempty" xorm:"-"`
	SPECS          []*GoodsSkuSpecMapping `json:"specs,omitempty" xorm:"-"`
}

type GoodsByPagination struct {
	Total   int64    `json:"total"`
	Pages   int      `json:"pages"`
	Page    int      `json:"page"`
	HasNext bool     `json:"has_next"`
	Content []*Goods `json:"content"`
}

type SafeGoods struct {
	sync.RWMutex
	ValueMap map[int64]*Goods //key : Goods id
}

var CacheSafeGoods *SafeGoods

func InitSafeGoods() (err error) {
	s := new(Goods)
	result, err := s.GetAllFormated()
	if err != nil {
		logger.Error(err)
		return
	}
	CacheSafeGoods = new(SafeGoods)
	CacheSafeGoods.ValueMap = make(map[int64]*Goods)
	CacheSafeGoods.RLock()
	defer CacheSafeGoods.RUnlock()
	for _, v := range result {
		CacheSafeGoods.ValueMap[v.Id] = v
	}
	return
}

type GoodsImages struct {
	Id    int64  `json:"id"`
	Gid   int64  `json:"g_id" xorm:"index"`
	Url   string `json:"url" xorm:"varchar(200)"`
	Thumb string `json:"thumb" xorm:"varchar(200)"`
}

func (self *Goods) GetAllFormated() (result []*Goods, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(Goods)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		err = formatGoods(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}
	return
}

func formatGoods(t *Goods) (err error) {
	if t.Id < 1 {
		return
	}
	if t.Stat != 0 {
		t.StatF, _ = GoodsStatMap[t.Stat]
	}
	gm := new(GoodsSkuSpecMapping)
	gm.GId = t.Id
	t.SPECS, err = gm.GetAllFormated()
	if err != nil {
		logger.Error(err)
		return
	}

	if len(t.SPECS) > 0 {
		var priceLow, priceHigh float32
		var i int

		//计算出最低价，最高价
		for _, v := range t.SPECS {
			if i == 0 {
				priceHigh = v.SIdF.Price
				priceLow = v.SIdF.Price
			} else {
				if v.SIdF.Price > priceHigh {
					priceHigh = v.SIdF.Price
				}
				if v.SIdF.Price < priceLow {
					priceLow = v.SIdF.Price
				}
			}
			i++
		}

		if priceLow == priceHigh {
			t.PriceF = formatMoney(priceLow)
		} else {
			t.PriceF = formatMoney(priceLow) + "~" + formatMoney(priceHigh)
		}
	} else {
		t.PriceF = formatMoney(t.Price)
	}

	gi := new(GoodsImages)
	gi.Gid = t.Id
	t.Images, err = gi.GetAll()
	if err != nil {
		logger.Error(err)
		return
	}
	return

}

func (self *Goods) GetAll() (result []*Goods, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(Goods)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return

}
func (self *Goods) GetByPM(startPageNo, limit int, order string, classifys []int64) (result *GoodsByPagination, err error) {
	result = new(GoodsByPagination)
	result.Page = startPageNo

	vdb1 := Db.OrderBy("id desc")
	if order != "" {
		vdb1.And(order)
	}
	if classifys != nil {
		vdb1.In("classify", classifys)
	}

	vdb2 := vdb1.Clone()

	result.Total, err = vdb1.Count(self)
	if err != nil {
		logger.Error(err)
		return
	}

	rows, err := vdb2.Cols("id").Limit(limit, (startPageNo-1)*limit).Rows(self)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		t := new(Goods)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		if t.Id != 0 {

			tt, ok := CacheSafeGoods.ValueMap[t.Id]
			if ok {
				result.Content = append(result.Content, tt)
			} else {
				err = fmt.Errorf("Goods id:%d not found in GoodsCache", t.Id)
				logger.Error(err)
				return
			}
		}
	}

	tp := int(result.Total) / limit
	if result.Total == 0 {
		result.Pages = 0
	} else if int(result.Total) <= limit {
		result.Pages = 1
	} else if tp%limit == 0 {
		result.Pages = tp
	} else {
		result.Pages = tp + 1
	}
	if result.Page < result.Pages {
		result.HasNext = true
	}

	return
}

func (self *Goods) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *Goods) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *Goods) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

func (self *Goods) Delete() (int64, error) {
	return Db.Id(self.Id).Delete(self)
}

func (self *GoodsImages) GetAll() (result []*GoodsImages, err error) {
	rows, err := Db.Rows(self)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		t := new(GoodsImages)
		err = rows.Scan(t)
		if err != nil {
			return
		}
		result = append(result, t)
	}
	return

}
func (self *GoodsImages) GetOne() (bool, error) {
	return Db.Get(self)
}

func (self *GoodsImages) Insert() (int64, error) {
	return Db.Insert(self)
}

func (self *GoodsImages) Update() (int64, error) {
	return Db.Id(self.Id).Update(self)
}

func (self *GoodsImages) Delete() (int64, error) {
	return Db.Id(self.Id).Delete(self)
}
