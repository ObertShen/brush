package model

type (
	// UserTag 对应user_tag表
	UserTag struct {
		ID        int64          `xorm:"BigInt(15) pk autoincr 'id'" json:"id"`
		Platform  string         `xorm:"varchar(16) notnull" json:"-"`
		UserID    int64          `xorm:"BigInt(15) notnull" json:"-"`
		UnionID   string         `xorm:"varchar(128) notnull" json:"-"`
		Tags      []nameAndValue `xorm:"json notnull" json:"-"`
		UpdatedAt int64          `xorm:"Int(11) notnull" json:"updatedAt"`
		CreatedAt int64          `xorm:"Int(11) notnull" json:"createdAt"`
	}

	// nameAndValue
	nameAndValue struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	// UserTagData UserTage 对象
	UserTagData struct {
		*DBConn
	}
)

var userTagDataIns *UserTagData

// GetUserTagDataIns 获取 UserTagData 的单例
func GetUserTagDataIns() *UserTagData {
	if userTagDataIns != nil {
		userTagDataIns = NewUserTagData()
	}

	return userTagDataIns
}

// NewUserTagData 创建 UserTagData 对象
func NewUserTagData() *UserTagData {
	utd := &UserTagData{GetConnIns()}
	utd.conn.Sync2(new(Weibo))
	return utd
}

// Get 获取单条记录
func (uta *UserTagData) Get(userTag *UserTag) (bool, error) {
	return uta.conn.Get(userTag)
}
