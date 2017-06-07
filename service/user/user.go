package user

import "brush/model"

type (
	dataAccess struct {
		user  *model.UserData
		weibo *model.WeiboUserData
		zhihu *model.ZhihuUserData
		tag   *model.UserTagData
	}

	// Service 供controller层调用
	Service struct {
		dataAccess *dataAccess
	}
)

// NewService 创建Service对象
func NewService() *Service {
	return &Service{&dataAccess{model.GetUserDataIns(), model.GetWeiboUserDataIns(), model.GetZhihuUserDataIns(), model.GetUserTagDataIns()}}
}

// ServiceIns Service的单例
var ServiceIns *Service

// GetServiceIns 获得service的单例
func GetServiceIns() *Service {
	if ServiceIns == nil {
		ServiceIns = NewService()
	}

	return ServiceIns
}

// GetWeiboUsersByKeyWord 获取指定条数的用户记录
func (us *Service) GetWeiboUsersByKeyWord(nickName string) ([]*Weibo, error) {
	records, err := us.dataAccess.weibo.GetByNickName(nickName)
	if err != nil {
		return nil, err
	}

	return us.getWeiboUsers(records)
}

// GetWeiboUsers 获取指定条数的用户记录
func (us *Service) GetWeiboUsers(weiboUser model.WeiboUser, pageSize, pageNo int) ([]*Weibo, error) {
	records, err := us.dataAccess.weibo.GetList(&weiboUser, pageSize, pageNo)
	if err != nil {
		return nil, err
	}

	return us.getWeiboUsers(records)
}

func (us *Service) getWeiboUsers(records []*model.WeiboUser) (weiboUsers []*Weibo, err error) {
	weiboUsers = []*Weibo{}
	for _, record := range records {
		weiboUsers = append(weiboUsers, &Weibo{*record, "weibo"})
	}

	return
}

// GetZhihuUsersByKeyWord 获取指定条数的用户记录
func (us *Service) GetZhihuUsersByKeyWord(keyWord string) ([]*Zhihu, error) {
	records, err := us.dataAccess.zhihu.GetByNickName(keyWord)
	if err != nil {
		return nil, err
	}

	recordsByUserName, err := us.dataAccess.zhihu.GetByUserName(keyWord)
	if err != nil {
		return nil, err
	}

	for _, recordByUserName := range recordsByUserName {
		notIn := false
		for _, record := range records {
			if recordByUserName.ID == record.ID {
				notIn = true
			}
		}

		if !notIn {
			records = append(records, recordByUserName)
		}
	}

	return us.getZhihuUsers(records)
}

// GetZhihuUsers 根据用户名获取知乎用户
func (us *Service) GetZhihuUsers(zhihuUser model.ZhihuUser, pageSize, pageNo int) (zhihuUsers []*Zhihu, err error) {
	records, err := us.dataAccess.zhihu.GetList(&zhihuUser, pageSize, pageNo)
	if err != nil {
		return nil, err
	}

	return us.getZhihuUsers(records)
}

func (us *Service) getZhihuUsers(records []*model.ZhihuUser) (zhihuUsers []*Zhihu, err error) {
	zhihuUsers = []*Zhihu{}
	for _, record := range records {
		zhihuUsers = append(zhihuUsers, &Zhihu{*record, "zhihu"})
	}

	return
}

// GetZhihuUserByFollower 获取某个知乎用户的关注者信息
func (us *Service) GetZhihuUserByFollower(zhihuID string) ([]*Node, []*Link, error) {
	nodes := []*Node{}
	links := []*Link{}

	zhihuUser := &model.ZhihuUser{UserName: zhihuID}
	has, err := us.dataAccess.zhihu.Get(zhihuUser)
	if err != nil {
		return nil, nil, err
	}

	if !has {
		return nodes, links, nil
	}

	records, err := us.dataAccess.zhihu.GetListByFollower(zhihuID)
	if err != nil {
		return nil, nil, err
	}

	nodes = append(nodes, &Node{Name: zhihuUser.NickName, Category: 0, Value: 10, Label: zhihuUser.NickName + "\n(主要)"})
	for _, record := range records {
		nodes = append(nodes, &Node{Category: 1, Name: record.NickName, Value: 8})
		links = append(links, &Link{Source: zhihuUser.NickName, Target: record.NickName, Weight: 5})
	}

	return nodes, links, nil
}

// GetWeiboUserByFollower 获取某个微博用户的关注者信息
func (us *Service) GetWeiboUserByFollower(weiboUserID int64) ([]*Node, []*Link, error) {
	nodes := []*Node{}
	links := []*Link{}

	weiboUser := &model.WeiboUser{ID: weiboUserID}
	has, err := us.dataAccess.weibo.Get(weiboUser)
	if err != nil {
		return nil, nil, err
	}

	if !has {
		return nodes, links, nil
	}

	records, err := us.dataAccess.weibo.GetListByFollower(weiboUserID)
	if err != nil {
		return nil, nil, err
	}

	nodes = append(nodes, &Node{Name: weiboUser.NickName, Category: 0, Value: 10, Label: weiboUser.NickName + "\n(主要)"})
	for _, record := range records {
		nodes = append(nodes, &Node{Category: 1, Name: record.NickName, Value: 8})
		links = append(links, &Link{Source: record.NickName, Target: weiboUser.NickName, Weight: 5})
	}

	return nodes, links, nil
}
