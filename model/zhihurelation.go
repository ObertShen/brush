package model

// ZhihuRelation 与数据库中的 zhihu_relation 表相对应 用于存储项目的巡查信息
type ZhihuRelation struct {
	ID        int64  `xorm:"Int(11) pk autoincr 'id'" json:"-"`
	Follower  string `xorm:"varchar(64) notnull" json:"follower"`
	Followee  string `xorm:"varchar(64) notnull" json:"followee"`
	UpdatedAt int64  `xorm:"Int(11) notnull" json:"-"`
	CreatedAt int64  `xorm:"Int(11) notnull" json:"-"`
}

var (
	// ZhihuRelationDataIns 是 ZhihuRelationData 类的单例
	ZhihuRelationDataIns *ZhihuRelationData
)

// ZhihuRelationData is a class that includeing fundamental functions for CURD operations in zhihu_relation table
type ZhihuRelationData struct {
	*DBConn
}

// GetZhihuRelationDataIns 用于获取 ZhihuRelationData 类的单例
func GetZhihuRelationDataIns() *ZhihuRelationData {
	if ZhihuRelationDataIns == nil {
		ZhihuRelationDataIns = NewZhihuRelationData()
	}

	return ZhihuRelationDataIns
}

// NewZhihuRelationData init ZhihuRelationData class
func NewZhihuRelationData() *ZhihuRelationData {
	zrd := &ZhihuRelationData{GetConnIns()}
	zrd.conn.Sync2(new(ZhihuRelation))
	return zrd
}

// Insert 插入一条数据
func (zrd *ZhihuRelationData) Insert(record *ZhihuRelation) (int64, error) {
	return zrd.conn.Insert(record)
}

// GetList 获取record中的非空字段多条数据
func (zrd *ZhihuRelationData) GetList(record *ZhihuRelation) (records []*ZhihuRelation, err error) {
	records = []*ZhihuRelation{}
	if err = zrd.conn.Find(&records, record); err != nil {
		return
	}

	return
}
