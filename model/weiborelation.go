package model

// WeiboRelation 与数据库中的 weibo_relation 表相对应 用于存储项目的巡查信息
type WeiboRelation struct {
	ID         int64 `xorm:"Int(11) pk autoincr 'id'" json:"id"`
	FollowID   int64 `xorm:"BigInt(11) unique(relation) notnull" json:"-"`
	FollowerID int64 `xorm:"BigInt(11) unique(relation) notnull" json:"-"`
	UpdatedAt  int64 `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt  int64 `xorm:"Int(11) notnull" json:"createdAt"`
}

var (
	// WeiboRelationDataIns 是 WeiboRelationData 类的单例
	WeiboRelationDataIns *WeiboRelationData
)

// WeiboRelationData is a class that includeing fundamental functions for CURD operations in weibo_relation table
type WeiboRelationData struct {
	*DBConn
}

// GetWeiboRelationDataIns 用于获取 WeiboRelationData 类的单例
func GetWeiboRelationDataIns() *WeiboRelationData {
	if WeiboRelationDataIns == nil {
		WeiboRelationDataIns = NewWeiboRelationData()
	}

	return WeiboRelationDataIns
}

// NewWeiboRelationData init WeiboRelationData class
func NewWeiboRelationData() *WeiboRelationData {
	wrd := &WeiboRelationData{GetConnIns()}
	wrd.conn.Sync2(new(WeiboRelation))
	return wrd
}

// Insert 插入一条数据
func (wrd *WeiboRelationData) Insert(record *WeiboRelation) (int64, error) {
	return wrd.conn.Insert(record)
}
