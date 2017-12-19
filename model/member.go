package model

const (
	RANK_0 = 0
)

//会员体系
type Member struct {
	ID                int64   `json:"id,omitempty"`
	Guid              string  `json:"Guid,omitempty" xorm:"varchar(100) index"`      //是	会员唯一标识
	CardId            string  `json:"CardId,omitempty" xorm:"varchar(100) index"`    //是	会员卡号
	MemberGroupName   string  `json:"MemberGroupName,omitempty" xorm:"varchar(100)"` //是	会员级别
	TrueName          string  `json:"TrueName,omitempty" xorm:"varchar(100)"`        //是	姓名
	Sex               string  `json:"Sex,omitempty" xorm:"varchar(10)"`              //是	性别（1：先生，2：女士）
	Mobile            string  `json:"Mobile,omitempty" xorm:"varchar(50)"`           //是	手机号码
	IdCard            string  `json:"IdCard,omitempty" xorm:"varchar(50)"`           //是	身份证号码
	Tel               string  `json:"Tel,omitempty" xorm:"varchar(50)"`              //是	电话
	Email             string  `json:"email,omitempty" xorm:"varchar(200)"`           //是	Email
	ImagePath         string  `json:"ImagePath,omitempty" xorm:"varchar(100)"`       //是	头像图片路径
	ChainStoreName    string  `json:"ChainStoreName,omitempty" xorm:"varchar(100)"`  //是	注册所在店面
	ChainStoreGuid    string  `json:"ChainStoreGuid,omitempty" xorm:"varchar(100)"`  //是	注册所在店面唯一标识
	UserAccount       string  `json:"UserAccount,omitempty" xorm:"varchar(100)"`     //是	注册当事工号
	RegisterTime      string  `json:"RegisterTime,omitempty" xorm:"varchar(100)"`    //是	注册时间
	DurationTime      string  `json:"DurationTime,omitempty" xorm:"varchar(100)"`    //是	有效期
	TotalPoint        float32 `json:"TotalPoint,omitempty" xorm:"decimal(10,2)"`     //是	历史积分
	AvailablePoint    float32 `json:"AvailablePoint,omitempty" xorm:"decimal(10,2)"` //是	可用积分
	TotalValue        float32 `json:"TotalValue,omitempty" xorm:"decimal(10,2)"`     //是	历史储值
	AvailableValue    float32 `json:"AvailableValue,omitempty" xorm:"decimal(10,2)"` //是	可用储值,只用用于消费，否则扣除相应积分
	FreezedValue      float32 `json:"FreezedValue,omitempty" xorm:"decimal(10,2)"`   //是	冻结储值
	Meno              string  `json:"Meno,omitempty" xorm:"varchar(200)"`            //是	备注
	RecommendCardId   string  `json:"RecommendCardId,omitempty" xorm:"varchar(100)"` //是	推荐人卡号
	IsLocked          bool    `json:"IsLocked,omitempty" xorm:""`                    //是	是否锁定
	IsDeleted         bool    `json:"IsDeleted,omitempty" xorm:""`                   //是	是否删除到回收站（无法筛选已经彻底删除的会员）
	IsLunar           int     `json:"IsLunar,omitempty" xorm:"tinyint(2)"`           //是	生日是否是阴历
	RealBirthDay      string  `json:"RealBirthDay,omitempty" xorm:"varchar(100)"`    //是	生日（如果是阴历，则格式为“壬申年(1992)三月十六”,否则格式为“2015-05-28”）
	Birthday          string  `json:"Birthday,omitempty" xorm:"varchar(100)"`        //是	生日（如果是阴历，此字段已经转换为阳历）
	ModifiedTime      string  `json:"ModifiedTime,omitempty" xorm:"varchar(100)"`    //是	最近编辑时间
	ModifiedUser      string  `json:"ModifiedUser,omitempty" xorm:"varchar(100)"`    //是	最近编辑者
	NickName          string  `json:"nick_name,omitempty" xorm:"varchar(100)"`
	City              string  `json:"city,omitempty" xorm:"varchar(100)"`
	Province          string  `json:"province,omitempty" xorm:"varchar(100)"`
	Country           string  `json:"country,omitempty" xorm:"varchar(100)"`
	WxAvatarPath      string  `json:"avatar_path,omitempty" xorm:"varchar(100)"`       //否	微信头像图片路径
	TotalAmount       float32 `json:"total_amount,omitempty" xorm:"decimal(10,2)"`     //总共赚的钱
	AvailableAmount   float32 `json:"available_amount,omitempty" xorm:"decimal(10,2)"` //还没有提走的钱
	ZoneID            int64   `json:"zone_id,omitempty" xorm:"index"`                  //隶属的区域会员ID
	RecommendMemberID int64   `json:"recommend_member_id,omitempty" xorm:"index"`      //上线ID-》谁推荐的
	ChainStoreId      int64   `json:"chain_store_id,omitempty" xorm:"index"`           //管理的配送店ID，仅有1级会员拥有
	OpenID            string  `json:"open_id,omitempty" xorm:"varchar(100) index"`
	UnionID           string  `json:"union_id,omitempty" xorm:"varchar(100)"`
}

func (self *Member) GetLastOne() (has bool, err error) {
	return Db.OrderBy("i_d desc").Limit(1).Get(self)
}
func (self *Member) GetOne() (has bool, err error) {
	return Db.Get(self)
}

func (self *Member) Insert() (num int64, err error) {

	return Db.Insert(self)
}

func (self *Member) Update() (num int64, err error) {
	return
}
