package model

import (
	"database/sql"

	"brush/core/db"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

var (
	// connIns 是 conn 的单例
	connIns *DBConn
	// weiboConnIns 是 conn 的单例
	weiboConnIns *DBConn
)

// DBConn 其他数据操作类的父类
type DBConn struct {
	conn  *xorm.Engine
	redis *redis.Client
}

// InitDB 初始化数据库
func InitDB() {
	GetZhihuConnIns()
	GetConnIns()

	GetCrawlerDataIns()
	GetUserDataIns()
	GetUserTagDataIns()
	GetWeiboDataIns()
	GetWeiboRelationDataIns()
	GetWeiboUserDataIns()
	GetZhihuRelationDataIns()
	GetZhihuUserDataIns()
}

// GetZhihuConnIns 用于获取 DBConn 类的单例
func GetZhihuConnIns() *DBConn {
	if connIns == nil {
		connIns = &DBConn{db.GetZhihuInstance().Conn, db.GetRedisClientIns()}
	}

	return connIns
}

// GetConnIns 用于获取 DBConn 类的单例
func GetConnIns() *DBConn {
	if weiboConnIns == nil {
		weiboConnIns = &DBConn{db.GetDefaultInstance().Conn, db.GetRedisClientIns()}
	}

	return weiboConnIns
}

// Query 用于简单查询 查询单个字段等
func (md *DBConn) Query(sql string, paramStr ...interface{}) (results []map[string][]byte, err error) {
	return md.conn.Query(sql, paramStr...)
}

// Exec 用于直接执行一个SQL命令
func (md *DBConn) Exec(sql string, args ...interface{}) (result sql.Result, err error) {
	return md.conn.Exec(sql, args...)
}
