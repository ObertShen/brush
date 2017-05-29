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

var (
	dbDebug = 0
)

const (
	maxIdleConns = 8
	maxOpenConns = 300
)

// Manager 数据库对象
type Manager struct {
	Conn       *xorm.Engine
	DebugIndex int
}

// NewManager 初始化数据库对象 提供xorm的方法
func NewManager() *Manager {
	mgr := new(Manager)
	mgr.DebugIndex = dbDebug
	dbDebug++
	if dbDebug >= 32768 {
		dbDebug = 0
	}

	return mgr
}

// CloseConnect 关闭ORM引擎 一般不用手动关闭 程序退出时自动关闭
func (dm *Manager) CloseConnect() {
	if err := dm.Conn.Close(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("[%d]Disconnect DB : %s \n", this.DebugIndex, dbType)
}

// OpenConnect 生成ORM引擎，建立数据库连接
func (dm *Manager) OpenConnect() {
	engine, err := xorm.NewEngine("mysql", config.GetConfig().GetValue("DBURL"))
	if err != nil {
		log.Printf("[%d]DB Open Failed : %s \n", dm.DebugIndex, "mysql")
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
