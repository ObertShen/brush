package model

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

type (
	// UserTag 对应user_tag表
	UserTag struct {
		ID        int64          `xorm:"BigInt(15) pk autoincr 'id'" json:"id"`
		Platform  string         `xorm:"varchar(16) notnull" json:"-"`
		UserID    int64          `xorm:"BigInt(15) notnull" json:"-"`
		Tags      []nameAndValue `xorm:"json notnull" json:"-"`
		UpdatedAt int64          `xorm:"Int(11) notnull" json:"updatedAt"`
		CreatedAt int64          `xorm:"Int(11) notnull" json:"createdAt"`
	}

	// nameAndValue
	nameAndValue struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	// UserTagData UserTage 对象
	UserTagData struct {
		*DBConn
	}
)

var userTagDataIns *UserTagData

// GetUserTagDataIns 获取 UserTagData 的单例
func GetUserTagDataIns() *UserTagData {
	if userTagDataIns == nil {
		userTagDataIns = NewUserTagData()
	}

	return userTagDataIns
}

// NewUserTagData 创建 UserTagData 对象
func NewUserTagData() *UserTagData {
	utd := &UserTagData{GetConnIns()}
	utd.conn.Sync2(new(UserTag))
	return utd
}

// Get 获取单条记录
func (uta *UserTagData) Get(userTag *UserTag) (bool, error) {
	return uta.conn.Get(userTag)
}

// GetList 获取多条记录
func (uta *UserTagData) GetList(userTag *UserTag) (tags []*UserTag, err error) {
	tags = []*UserTag{}
	if err = uta.conn.Find(&tags, userTag); err != nil {
		return
	}

	return
}

// GetWeiboTags 获取微博tag
func (uta *UserTagData) GetWeiboTags(weiboUserID int64, nickName string) (val string, err error) {
	fmt.Println("weibo-" + strconv.FormatInt(weiboUserID, 10) + "-" + nickName)
	val, err = uta.redis.Get("weibo-" + strconv.FormatInt(weiboUserID, 10) + nickName).Result()
	if err == redis.Nil {
		return "", nil
	}

	return
}

// GetZhihuTags 获取zhihutag
func (uta *UserTagData) GetZhihuTags(zhihuID int64, nickName string) (val string, err error) {
	fmt.Println("zhihu-" + strconv.FormatInt(zhihuID, 10) + nickName)
	val, err = uta.redis.Get("zhihu-" + strconv.FormatInt(zhihuID, 10) + "-" + nickName).Result()
	if err == redis.Nil {
		return "", nil
	}

	return
}

// Insert 插入单条记录
func (uta *UserTagData) Insert(userTag *UserTag) (int64, error) {
	return uta.conn.Insert(userTag)
}
