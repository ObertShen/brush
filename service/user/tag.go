package user

import (
	"brush/model"
	"encoding/json"
)

func (us *Service) GetWeiboUserTags(weiboUserID int64) (userTag *model.UserTag, err error) {
	userTag = &model.UserTag{UserID: weiboUserID, Platform: "weibo"}
	has, err := us.dataAccess.tag.Get(userTag)
	if err != nil || !has {
		return nil, err
	}

	return
}

func (us *Service) GetZhihuUserTags(zhihuUserID string) (userTag *model.UserTag, err error) {
	zhihuUser := &model.ZhihuUser{UserName: zhihuUserID}
	has, err := us.dataAccess.zhihu.Get(zhihuUser)
	if err != nil {
		return
	}

	if !has {
		return nil, nil
	}

	userTag = &model.UserTag{UserID: zhihuUser.ID, Platform: "zhihu"}
	has, err = us.dataAccess.tag.Get(userTag)
	if err != nil || !has {
		return nil, err
	}

	return
}

// GetWeiboTagsFromRedis 从tags中获取tags
func (us *Service) GetWeiboTagsFromRedis(weiboID int64) (interface{}, error) {
	weiboUser := &model.WeiboUser{ID: weiboID}
	has, err := us.dataAccess.weibo.Get(weiboUser)
	if err != nil || !has {
		return nil, err
	}

	tags, err := us.dataAccess.tag.GetWeiboTags(weiboID, weiboUser.NickName)
	if err != nil {
		return nil, err
	}

	if tags == "" {
		return []string{}, nil
	}

	result := make(map[string]interface{})
	if err = json.Unmarshal([]byte(tags), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetZhihuTagsFromRedis 从tags中获取tags
func (us *Service) GetZhihuTagsFromRedis(zhihuID int64) (interface{}, error) {
	zhihuUser := &model.ZhihuUser{ID: zhihuID}
	has, err := us.dataAccess.zhihu.Get(zhihuUser)
	if err != nil || !has {
		return nil, err
	}

	tags, err := us.dataAccess.tag.GetZhihuTags(zhihuID, zhihuUser.NickName)
	if err != nil {
		return nil, err
	}

	if tags == "" {
		return []string{}, nil
	}

	result := make(map[string]interface{})
	if err = json.Unmarshal([]byte(tags), &result); err != nil {
		return nil, err
	}

	return result, nil
}
