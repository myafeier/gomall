package model

import (
	"errors"
)

type Order struct {
	ID                   int64        `json:"id,omitempty"`
	Guid                 string       `json:"Guid,omitempty" xorm:"varchar(100) index"`
	BillNumber           string       `json:"BillNumber,omitempty" xorm:"varchar(100) index"`     //是	订单号
	Status               int          `json:"Status,omitempty" xorm:"varchar(100) index"`         //是	订单状态(-1:订单取消,0:等待审核,1:审核通过,2:等待支付,3:支付成功，其他：不通过)
	CardId               string       `json:"CardId,omitempty" xorm:"varchar(100)"`               //是	卡号
	TrueName             string       `json:"TrueName,omitempty" xorm:"varchar(100)"`             //是	姓名
	Mobile               string       `json:"Mobile,omitempty" xorm:"varchar(100)"`               //是	手机号
	AssignChainStoreGuid string       `json:"AssignChainStoreGuid,omitempty" xorm:"varchar(100)"` //否	指派店面唯一标识
	AssignStoreName      string       `json:"AssignStoreName,omitempty" xorm:"varchar(100)"`      //否	指派店面名称
	StatusName           string       `json:"StatusName,omitempty" xorm:"varchar(100)"`           //是	订单状态名称
	SubmitTime           string       `json:"SubmitTime,omitempty" xorm:"varchar(100)"`           //是	提交时间
	TotalMoney           float32      `json:"TotalMoney,omitempty" xorm:"decimal(10,2)"`          //是	订单总金额
	PaidMoney            float32      `json:"PaidMoney,omitempty" xorm:"decimal(10,2)"`           //是	现金支付金额
	PaidValue            float32      `json:"PaidValue,omitempty" xorm:"decimal(10,2)"`           //是	储值支付金额
	PaidPoint            float32      `json:"PaidPoint,omitempty" xorm:"decimal(10,2)"`           //是	积分支付金额
	PaidCoupon           float32      `json:"PaidCoupon,omitempty" xorm:"decimal(10,2)"`          //是	优惠券支付金额
	PaidOther            float32      `json:"PaidOther,omitempty" xorm:"decimal(10,2)"`           //是	其他支付金额
	PaidThirdpay         float32      `json:"PaidThirdpay,omitempty" xorm:"decimal(10,2)"`        //是	第三方支付金额
	ThirdpayType         int          `json:"ThirdpayType,omitempty" xorm:"tinyint(2)"`           //是	第三方支付方式（1、微信支付，2、支付宝支付，0：非第三方支付）
	ConsumeBillNumber    string       `json:"ConsumeBillNumber,omitempty" xorm:"varchar(100)"`    //是	审核通过后的单据号（可在“获取消费列表”中查到）
	IsCashOnDeliver      bool         `json:"IsCashOnDeliver,omitempty" xorm:""`                  //是	支付方式（true：货到付款，false:在线支付 ）
	IsSelfPickUp         bool         `json:"IsSelfPickUp,omitempty" xorm:""`                     //是	配送方式（true ：上门自提，false：快递配送）
	SelfPickUpTime       string       `json:"SelfPickUpTime,omitempty" xorm:"varchar(100)"`       //是	自提时间
	SelfPickUpStoreName  string       `json:"SelfPickUpStoreName,omitempty" xorm:"varchar(100)"`  //是	自提门店
	Province             string       `json:"Province,omitempty" xorm:"varchar(100)"`             //是	送货地址-省
	City                 string       `json:"City,omitempty" xorm:"varchar(100)"`                 //是	送货地址-市
	County               string       `json:"County,omitempty" xorm:"varchar(100)"`               //是	送货地址-区
	Address              string       `json:"Address,omitempty" xorm:"varchar(100)"`              //是	送货地址-详细地址
	Postcode             string       `json:"Postcode,omitempty" xorm:"varchar(100)"`             //是	邮编
	Receiver             string       `json:"Receiver,omitempty" xorm:"varchar(100)"`             //是	收货人（为空时参考TrueName）
	ReceiverMobile       string       `json:"ReceiverMobile,omitempty" xorm:"varchar(100)"`       //是	联系方式（为空时参考Mobile）
	Meno                 string       `json:"Meno,omitempty" xorm:"varchar(200)"`                 //是	备注
	ItemList             []*OrderItem `json:"ItemList,omitempty" xorm:"-"`                        //是	订单明细（BarCode:商品编码,Name:商品名称,Number:数量,Price:单价）
	//DispatchUid          int64        `json:"dispatch_uid" xorm:"index"`
	//DispatchTime         time.Time    `json:"dispatch_time"`
}

//订单明细
type OrderItem struct {
	ID      int64  `json:"id"`
	OrderID int64  `json:"order_id" xorm:"index"`                 //订单ID
	BarCode string `json:"BarCode,omitempty" xorm:"varchar(100)"` //商品编码,
	Name    string `json:"Name,omitempty" xorm:"varchar(100)"`    //:商品名称,
	Number  string `json:"Number,omitempty" xorm:"varchar(100)"`  //:数量,
	Price   string `json:"Price,omitempty" xorm:"varchar(100)"`   //:单价
}

func (self *Order) GetOneWithItems() (has bool, err error) {
	has, err = Db.Get(self)
	if err != nil {
		logger.Error(err)
		return
	}
	if has {
		oi := new(OrderItem)
		oi.OrderID = self.ID
		self.ItemList, err = oi.GetList()
		if err != nil {
			logger.Error(err)
			return
		}
	}
	return
}

func (self *Order) GetLastOne() (has bool, err error) {
	return Db.OrderBy("i_d desc").Limit(1).Get(self)
}

func (self *Order) GetOne() (has bool, err error) {
	return Db.Get(self)
}
func (self *Order) Insert() (num int64, err error) {
	return Db.Insert(self)
}
func (self *Order) Update() (num int64, err error) {
	return Db.Id(self.ID).Update(self)
}

func (self *Order) RefreshOrderItem() error {
	if self.ID == 0 {
		return errors.New("order ID is 0")
	}
	for _, v := range self.ItemList {
		v.OrderID = self.ID
		has, err := v.GetOne()
		if err != nil {
			return err
		}

		if !has {
			_, err = v.Insert()
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (self *OrderItem) GetOne() (has bool, err error) {
	return Db.Get(self)
}

func (self *OrderItem) Insert() (num int64, err error) {
	return Db.Insert(self)
}
func (self *OrderItem) Update() (num int64, err error) {
	return Db.Id(self.ID).Update(self)
}
func (self *OrderItem) GetList() (result []*OrderItem, err error) {
	rows, err := Db.OrderBy("i_d").Rows(self)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		t := new(OrderItem)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}

	return
}
