package user

import "brush/model"

// NewService 创建Service对象
func NewService() *Service {
	return &Service{&dataAccess{model.GetUserDataIns(), model.GetWeiboUserDataIns(), model.GetZhihuUserDataIns()}}
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

// GetZhihuUserListByName 根据用户名获取知乎用户
func (us *Service) GetZhihuUserListByName(nickName string) (zhihuUsers []*Zhihu, err error) {
	records, err := us.dataAccess.zhihu.GetList(&model.ZhihuUser{NickName: nickName})
	if err != nil {
		return nil, err
	}

	zhihuUsers = []*Zhihu{}
	for _, record := range records {
		zhihuUsers = append(zhihuUsers, &Zhihu{*record, "zhihu"})
	}

	return
}

// GetWeiboUserListByName 根据用户名获取微博用户
func (us *Service) GetWeiboUserListByName(nickName string) (weiboUsers []*Weibo, err error) {
	records, err := us.dataAccess.weibo.GetList(&model.WeiboUser{NickName: nickName})
	if err != nil {
		return nil, err
	}

	weiboUsers = []*Weibo{}
	for _, record := range records {
		weiboUsers = append(weiboUsers, &Weibo{*record, "weibo"})
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
