package model

type ChainStore struct {
	ID int64
	Guid string 	//是	店面唯一标识
	StoreName string //	是	店面名称
	Contact	string //是	联系人
	Tel	string //是	电话
	EnablePointLimit	int64 //是	是否开启积分受限
	AvailableLimitPoint	int64 //是	可用积分受限额
	EnableValueLimit	int64 //是	是否开启储值受限
	AvailableLimitValue	int64 //是	可用储值受限额
	EnableSmsLimit	int64  //是	是否开启短信受限
	AvailableLimitSms	int64 //是	可用短信受限数量
	ProvinceName	string //是	省份
	CityName	string //是	城市
	CountyName	string //是	区县
	Address	string //是	详细地址
	ImagePath	string //是	图片路径
	Description	string //是	描述
	Meno	string //是	备注

	Stat    int  // 状态，自定义
	AdminId int64 //管理员ID，自定义
}
