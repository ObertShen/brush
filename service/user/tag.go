package user

import "brush/model"

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
