package user

import "brush/model"

type (
	dataAccess struct {
		user  *model.UserData
		weibo *model.WeiboUserData
		zhihu *model.ZhihuUserData
	}
	// Service 供controller层调用
	Service struct {
		dataAccess *dataAccess
	}
)

// NewService 创建新
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

// GetZhihuRelations 获取人物关系
// func (us *Service) GetZhihuRelations(zhihu string) ([]*Node, []*Link, error) {
// 	userList, err := us.GetList(&model.ZhihuUser{UserName: userName})
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return nil, nil, nil
// }

// GetWeiboRelations 获取人物关系
func (us *Service) GetWeiboRelations(weiboUserID int64) ([]*Node, []*Link) {

	return nil, nil
}
