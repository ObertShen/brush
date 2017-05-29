package model

import "log"

// User 与数据库中的 project_inspection 表相对应 用于存储项目的巡查信息
type User struct {
	ID        int64  `xorm:"Int(11) pk autoincr 'id'" json:"-"`
	UnionID   string `xorm:"varchar(128) unique notnull" json:"-"`
	Name      string `xorm:"varchar(128) unique notnull" json:"username"`
	NickName  string `xorm:"varchar(128) notnull" json:"nickname"`
	Password  string `xorm:"varchar(128) notnull" json:"-"`
	Role      string `xorm:"Enum('admin','visitor') notnull" json:"role"`
	Icon      string `xorm:"varchar(512) notnull" json:"icon"`
	UpdatedAt int64  `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt int64  `xorm:"Int(11) notnull" json:"createdAt"`
}

var (
	// UserDataIns 是 UserData 类的单例
	UserDataIns *UserData
)

// UserData is a class that includeing fundamental functions for CURD operations in user table
type UserData struct {
	*DBConn
}

// GetUserDataIns 用于获取 UserData 类的单例
func GetUserDataIns() *UserData {
	if UserDataIns == nil {
		UserDataIns = NewUserData()
	}

	return UserDataIns
}

// NewUserData init UserData class
func NewUserData() *UserData {
	ud := &UserData{GetConnIns()}
	if err := ud.conn.Sync2(new(User)); err != nil {
		log.Panic(err)
	}
	return ud
}

// GetList 从 user 中查询纪录
func (ud *UserData) GetList(user *User) (userList []*User, err error) {
	userList = []*User{}
	if err = ud.conn.UseBool("is_deleted", "is_admin").Find(&userList, user); err != nil {
		return
	}

	return
}

// Get 从 user 中查询纪录
func (ud *UserData) Get(user *User) (bool, error) {
	return ud.conn.UseBool("is_deleted", "is_admin").Get(user)
}

// Insert 从 project_inspection 中增加纪录
func (ud *UserData) Insert(user *User) (int64, error) {
	return ud.conn.Insert(user)
}
