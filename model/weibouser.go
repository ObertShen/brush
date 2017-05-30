package model

// WeiboUser 与数据库中的 weibo_user 表相对应 用于存储项目的巡查信息
type WeiboUser struct {
	ID           int64  `xorm:"BigInt(15) pk autoincr 'id'" json:"id"`
	FollowNumber int64  `xorm:"Int(11) notnull" json:"followings"`
	FansNumber   int64  `xorm:"Int(11) notnull" json:"followers"`
	WeiboNumber  int64  `xorm:"notnull" json:"-"`
	NickName     string `xorm:"varchar(128) notnull" json:"username"`
	Region       string `xorm:"varchar(128) notnull" json:"region"`
	Gender       int    `xorm:"TinyInt(2) notnull" json:"gender"`
	Birthday     string `xorm:"varchar(64) notnull" json:"birthday"`
	Introduction string `xorm:"Text notnull" json:"introduction"`
	RegisterTime string `xorm:"varchar(64) notnull" json:"-"`
	UpdatedAt    int64  `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt    int64  `xorm:"Int(11) notnull" json:"createdAt"`
}

var (
	// WeiboUserDataIns 是 WeiboUserData 类的单例
	WeiboUserDataIns *WeiboUserData
)

// WeiboUserData is a class that includeing fundamental functions for CURD operations in weibo_user table
type WeiboUserData struct {
	*DBConn
}

// GetWeiboUserDataIns 用于获取 WeiboUserData 类的单例
func GetWeiboUserDataIns() *WeiboUserData {
	if WeiboUserDataIns == nil {
		WeiboUserDataIns = NewWeiboUserData()
	}

	return WeiboUserDataIns
}

// NewWeiboUserData init UserData class
func NewWeiboUserData() *WeiboUserData {
	ud := &WeiboUserData{GetConnIns()}
	ud.conn.Sync2(new(WeiboUser))
	return ud
}

// Insert 插入一条数据
func (wud *WeiboUserData) Insert(record *WeiboUser) (int64, error) {
	return wud.conn.Insert(record)
}

// GetList 从 weibo_user 中查询纪录
func (wud *WeiboUserData) GetList(record *WeiboUser) (userList []*WeiboUser, err error) {
	userList = []*WeiboUser{}
	if err = wud.conn.UseBool("is_deleted").Find(&userList, record); err != nil {
		return
	}

	return
}

// Get 从 weibo_user 中查询纪录
func (wud *WeiboUserData) Get(record *WeiboUser) (bool, error) {
	return wud.conn.UseBool("is_deleted").Get(record)
}
