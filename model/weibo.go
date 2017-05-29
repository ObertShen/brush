package model

// Weibo 与数据库中的 Weibo 表相对应 用于存储项目的巡查信息
type Weibo struct {
	ID            int64  `xorm:"BigInt(15) pk autoincr 'id'" json:"id"`
	UserID        int64  `xorm:"BigInt(15)" json:"-"`
	CommentNumber int64  `xorm:"Int(11) notnull" json:"-"`
	LikeNumber    int64  `xorm:"Int(11) notnull" json:"-"`
	SentTime      int64  `xorm:"BigInt(15) notnull" json:"-"`
	Content       string `xorm:"varchar(512) notnull" json:"-"`
	UpdatedAt     int64  `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt     int64  `xorm:"Int(11) notnull" json:"createdAt"`
}

var (
	// WeiboDataIns 是 WeiboData 类的单例
	WeiboDataIns *WeiboData
)

// WeiboData is a class that includeing fundamental functions for CURD operations in weibo_user table
type WeiboData struct {
	*DBConn
}

// GetWeiboDataIns 用于获取 WeiboData 类的单例
func GetWeiboDataIns() *WeiboData {
	if WeiboDataIns == nil {
		WeiboDataIns = NewWeiboData()
	}

	return WeiboDataIns
}

// NewWeiboData init WeiboData class
func NewWeiboData() *WeiboData {
	ud := &WeiboData{GetConnIns()}
	ud.conn.Sync2(new(Weibo))
	return ud
}

// Insert 插入一条数据
func (wd *WeiboData) Insert(record *Weibo) (int64, error) {
	return wd.conn.Insert(record)
}
