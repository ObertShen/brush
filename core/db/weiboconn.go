package db

import (
	"log"
	"os"

	"brush/core/config"

	// Init the mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

const (
	maxIdleConns = 8
	maxOpenConns = 300
)

// WeiboManager 数据库对象
type WeiboManager struct {
	Conn *xorm.Engine
}

// NewWeiboManager 初始化数据库对象 提供xorm的方法
func NewWeiboManager() *WeiboManager {
	mgr := new(WeiboManager)
	return mgr
}

// CloseConnect 关闭连接池 一般不用手动关闭 程序退出时自动关闭
func (dm *WeiboManager) CloseConnect() {
	if err := dm.Conn.Close(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("[%d]Disconnect DB : %s \n", this.DebugIndex, dbType)
}

// OpenConnect 生成ORM引擎，建立数据库连接
func (dm *WeiboManager) OpenConnect() {
	engine, err := xorm.NewEngine("mysql", config.GetConfig().GetValue("DefaultDBURL"))
	if err != nil {
		log.Printf("DB Open Failed : %s \n", "mysql_weibo")
		log.Fatal(err)
	}

	engine.SetMaxIdleConns(maxIdleConns)
	engine.SetMaxOpenConns(maxOpenConns)
	engine.SetMapper(core.GonicMapper{})
	// 如果为测试环境 则会打出SQL语句和警告
	if os.Getenv("GIN_MODE") != "release" {
		engine.Logger().SetLevel(core.LOG_DEBUG)
		engine.ShowSQL(true)
	} else {
		engine.Logger().SetLevel(core.LOG_ERR)
	}
	dm.Conn = engine
	//fmt.Printf("[%d]Connected DB : %s \n", this.DebugIndex, dbType)
}
