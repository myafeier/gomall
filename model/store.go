package model

type Store struct {
	Id       int64  `json:"id"`
	Guid     string `json:"guid" xorm:"varchar(50) index"`
	Name     string `json:"name" xorm:"varchar(50) index"`
	Contact  string `json:"contact" xorm:"varchar(50)"` //联系人
	Tel      string `json:"tel" xorm:"varchar(50)"`
	Province string `json:"province" xorm:"varchar(50)"`
	City     string `json:"city" xorm:"varchar(50)"`
	County   string `json:"county" xorm:"varchar(50)"`
	Address  string `json:"address" xorm:"varchar(50)"`
}

func (self *Store) GetOne() (has bool, err error) {
	return Db.Get(self)
}

func (self *Store) Insert() (num int64, err error) {
	return Db.Insert(self)
}

func (self *Store) Update() (num int64, err error) {
	return Db.Id(self.Id).Update(self)
}
