package model

import (
	"time"
)

type Banner struct {
	Title     string    `json:"title" xorm:"varchar(200)"`  //: String,
	Remark    string    `json:"remark" xorm:"varchar(250)"` //: String,
	Sort      int       `json:"sort" xorm:"tinyint(2)"`     //: Number,
	IsShow    bool      `json:"is_show" xorm:"index"`       //: Boolean,
	Images    string    `json:"images" xorm:"varchar(250)"` //: Array,
	CreateAt  time.Time `json:"create_at" xorm:""`          //:
	UpdateAt  time.Time `json:"update_at" xorm:""`          //: Date,
	CreateAtF string    `json:"create_at_f" xorm:"-"`
	UpdateAtF string    `json:"update_at_f" xorm:"-"`
}

func (self *Banner) GetAll() (result []*Banner, err error) {

	rows, err := Db.Rows(self)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		t := new(Banner)
		err = rows.Scan(t)
		if err != nil {
			logger.Error(err)
			return
		}
		result = append(result, t)
	}
	return
}
