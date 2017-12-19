package model

import (
	//"common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

//数据库核心
var Db *xorm.Engine

//初始化数据库
func InitDb(dbType, host, port, dbuser, dbpass, dbname string) *xorm.Engine {
	var err error
	Db, err = xorm.NewEngine(dbType, dbuser+":"+dbpass+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8mb4")
	if err != nil {
		panic("error when connect database!,err:" + err.Error())
	}
	Db.SetMaxIdleConns(10)
	Db.SetMaxOpenConns(100)
	//Db.NoAutoCondition(true)

	//nMapper := core.NewPrefixMapper(core.SnakeMapper{}, "wss_")
	//Db.SetTableMapper(nMapper)
	Db.ShowSQL(true)

	//初始化客户来源数据库

	//测试环境数据库初始化开关
	isExist := false

	isExist, err = Db.IsTableExist(&Member{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Member{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Member{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&Store{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Store{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Store{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&User{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&User{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&User{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&UserAccountLog{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&UserAccountLog{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&UserAccountLog{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}
	isExist, err = Db.IsTableExist(&Order{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Order{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Order{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&OrderItem{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&OrderItem{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&OrderItem{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&OrderDispatchMission{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&OrderDispatchMission{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&OrderDispatchMission{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&Banner{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Banner{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Banner{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&Classify{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Classify{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Classify{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}
	isExist, err = Db.IsTableExist(&UserRankClassifyDiscount{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&UserRankClassifyDiscount{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&UserRankClassifyDiscount{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&GoodsSpecAttr{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&GoodsSpecAttr{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&GoodsSpecAttr{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&GoodsSpecValue{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&GoodsSpecValue{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&GoodsSpecValue{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&GoodsSku{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&GoodsSku{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&GoodsSku{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&GoodsSkuSpecMapping{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&GoodsSkuSpecMapping{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&GoodsSkuSpecMapping{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&GoodsImages{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&GoodsImages{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&GoodsImages{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	isExist, err = Db.IsTableExist(&Goods{})
	if err != nil {
		logger.Error("Check table existed error:", err)
	}
	if !isExist {
		err = Db.CreateTables(&Goods{})
		if err != nil {
			log.Fatal("create table error:", err)
		}
		err = Db.CreateIndexes(&Goods{})
		if err != nil {
			log.Fatal("create table index error:", err)
		}
	}

	err = InitClassify()
	if err != nil {
		log.Fatal("init classify error:", err)
	}
	err = InitUserRankClassifyDiscount()
	if err != nil {
		log.Fatal("init error:", err)
	}
	err = InitCacheSafeGoodsSpecAttr()
	if err != nil {
		log.Fatal("init error:", err)
	}
	err = InitCacheSafeGoodsSpecValue()
	if err != nil {
		log.Fatal("init error:", err)
	}

	//一定要先初始化sku
	err = InitSafeGoodsSku()
	if err != nil {
		log.Fatal("init error:", err)
	}
	//然后在初始化goods
	err = InitSafeGoods()
	if err != nil {
		log.Fatal("init error:", err)
	}

	return Db

}
