package db

// Manager 类的单例
var dbInstance *Manager

// GetInstance 获取数据库对象的单例
func GetInstance() *Manager {
	if dbInstance == nil {
		dbInstance = NewManager()
	}
	return dbInstance
}
