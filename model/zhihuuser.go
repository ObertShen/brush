package model

// ZhihuUser 与数据库中的 zhihu_user 表相对应 用于存储项目的巡查信息
type ZhihuUser struct {
	ID                   int64  `xorm:"BigInt(15) pk autoincr 'id'" json:"id"`
	ZhihuID              string `xorm:"varchar(128) notnull" json:"zhihuID"`
	UserName             string `xorm:"varchar(128) notnull 'username'" json:"username"`
	NickName             string `xorm:"varchar(128) notnull 'nickname'" json:"nickname"`
	AvatarURL            string `xorm:"varchar(512) notnull" json:"avatarURL"`
	Gender               int    `xorm:"TinyInt(2) notnull" json:"gender"`
	City                 string `xorm:"varchar(32) notnull" json:"city"`
	HeadLine             string `xorm:"varchar(512) notnull" json:"headline"`
	Industry             string `xorm:"varchar(32) notnull" json:"industry"`
	Company              string `xorm:"varchar(128) notnull" json:"company"`
	School               string `xorm:"varchar(64) notnull" json:"school"`
	Major                string `xorm:"varchar(64) notnull" json:"major"`
	FollowerNum          int64  `xorm:"Int(11) notnull" json:"followerNum"`
	FollowingNum         int64  `xorm:"Int(11) notnull" json:"followingNum"`
	ThanksNum            int64  `xorm:"Int(11) notnull" json:"thanksNum"`
	AgreeNum             int64  `xorm:"Int(11) notnull" json:"agreeNum"`
	FavorateNum          int64  `xorm:"Int(11) notnull" json:"favoriteNum"`
	FollowingQuestionNum int64  `xorm:"Int(11) notnull" json:"followingQuestionNum"`
	FollowingTopicNum    int64  `xorm:"Int(11) notnull" json:"followingTopicNum"`
	ArticleNum           int64  `xorm:"Int(11) notnull" json:"articleNum"`
	AnswerNum            int64  `xorm:"Int(11) notnull" json:"answerNum"`
	Introduction         string `xorm:"varchar(2048) notnull" json:"introduction"`
	SpiderAnswers        int64  `xorm:"Int(11) notnull" json:"spiderAnswers"`
	SpiderFollowing      int64  `xorm:"Int(11) notnull" json:"spiderFollowing"`
	SpiderFollower       int64  `xorm:"Int(11) notnull" json:"spiderFollower"`
	Status               int    `xorm:"TinyInt(2) notnull" json:"status"`
	UpdatedAt            int64  `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt            int64  `xorm:"Int(11) notnull" json:"createdAt"`
}

var (
	// ZhihuUserDataIns 是 ZhihuUserData 类的单例
	ZhihuUserDataIns *ZhihuUserData
)

// ZhihuUserData is a class that includeing fundamental functions for CURD operations in zhihu_user table
type ZhihuUserData struct {
	*DBConn
}

// GetZhihuUserDataIns 用于获取 ZhihuUserData 类的单例
func GetZhihuUserDataIns() *ZhihuUserData {
	if ZhihuUserDataIns == nil {
		ZhihuUserDataIns = NewZhihuUserData()
	}

	return ZhihuUserDataIns
}

// NewZhihuUserData init ZhihuUserData class
func NewZhihuUserData() *ZhihuUserData {
	zud := &ZhihuUserData{GetConnIns()}
	zud.conn.Sync2(new(ZhihuUser))
	return zud
}

// Insert 插入一条数据
func (zud *ZhihuUserData) Insert(record *ZhihuUser) (int64, error) {
	return zud.conn.Insert(record)
}

// GetByNickName 根据 nickname进行模糊查询
func (zud *ZhihuUserData) GetByNickName(nickName string, pageSize, pageNo int) (userList []*ZhihuUser, err error) {
	userList = []*ZhihuUser{}
	if nickName == "" {
		return
	}

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 50 {
		pageSize = 50
	}

	if err = zud.conn.UseBool("is_deleted").Where("nickname like ?", nickName+"%").Limit(pageSize, pageSize*(pageNo-1)).Find(&userList); err != nil {
		return
	}

	return
}

// GetByUserName 根据 username进行模糊查询
func (zud *ZhihuUserData) GetByUserName(userName string, pageSize, pageNo int) (userList []*ZhihuUser, err error) {
	userList = []*ZhihuUser{}
	if userName == "" {
		return
	}

	if pageNo < 1 {
		pageNo = 1
	}

	if pageSize < 50 {
		pageSize = 50
	}

	if err = zud.conn.UseBool("is_deleted").Where("username like ?", userName+"%").Limit(pageSize, pageSize*(pageNo-1)).Find(&userList); err != nil {
		return
	}

	return
}

// GetList 从 zhihu_user 中查询纪录
func (zud *ZhihuUserData) GetList(record *ZhihuUser, pageSize, pageNo int) (userList []*ZhihuUser, err error) {
	userList = []*ZhihuUser{}
	if err = zud.conn.UseBool("is_deleted").Limit(pageSize, pageSize*(pageNo-1)).Find(&userList, record); err != nil {
		return
	}

	return
}

// GetListByFollower 从 zhihu_user 中查询纪录
func (zud *ZhihuUserData) GetListByFollower(zhihuID string) (userList []*ZhihuUser, err error) {
	userList = []*ZhihuUser{}
	if err = zud.conn.Join("INNER", "zhihu_relation", "zhihu_relation.followee=zhihu_user.username AND zhihu_relation.follower=?", zhihuID).Find(&userList); err != nil {
		return
	}

	return
}

// Get 从 zhihu_user 中查询纪录
func (zud *ZhihuUserData) Get(record *ZhihuUser) (bool, error) {
	return zud.conn.UseBool("is_deleted").Get(record)
}
