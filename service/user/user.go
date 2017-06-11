package user

import "brush/model"
import "fmt"

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
func (us *Service) GetWeiboUsersByKeyWord(nickName string, pageSize, pageNo int) ([]*Weibo, error) {
	records, err := us.dataAccess.weibo.GetByNickName(nickName, pageSize, pageNo)
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

// GetZhihuUsersByKeyWord 根据关键字获取指定条数的用户记录
func (us *Service) GetZhihuUsersByKeyWord(keyWord string, pageSize, pageNo int) ([]*Zhihu, error) {
	records, err := us.dataAccess.zhihu.GetByNickName(keyWord, pageSize, pageNo)
	if err != nil {
		return nil, err
	}

	recordsByUserName, err := us.dataAccess.zhihu.GetByUserName(keyWord, pageSize, pageNo)
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

	nodes = append(nodes, &Node{ID: zhihuUser.ID, ZhihuID: zhihuUser.UserName, Name: zhihuUser.NickName, Category: 0, Value: 10, Label: zhihuUser.NickName + "\n(主要)"})
	for i, record := range records {
		if i >= 10 {
			return nodes, links, nil
		}

		nodes = append(nodes, &Node{ID: record.ID, ZhihuID: record.UserName, Category: 1, Name: record.NickName, Value: 9})
		links = append(links, &Link{Source: zhihuUser.NickName, Target: record.NickName, Weight: 9})
		nodes, links, err = us.getZhihuRelationCircle(nodes[0], nodes[i+1], nodes, links, 2)
		if err != nil {
			return nil, nil, err
		}
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

	nodes = append(nodes, &Node{ID: weiboUser.ID, Name: weiboUser.NickName, Category: 0, Value: 10, Label: weiboUser.NickName + "\n(主要)"})
	for i, record := range records {
		if i >= 10 {
			return nodes, links, nil
		}

		nodes = append(nodes, &Node{ID: record.ID, Category: 1, Name: record.NickName, Value: 9})
		links = append(links, &Link{Source: record.NickName, Target: weiboUser.NickName, Weight: 9})
		nodes, links, err = us.getRelationCircle(nodes[0], nodes[i+1], nodes, links, 2)
		if err != nil {
			return nil, nil, err
		}
	}

	return nodes, links, nil
}

func (us *Service) getRelationCircle(mainNode *Node, sourceNode *Node, nodes []*Node, links []*Link, times int) ([]*Node, []*Link, error) {
	records, err := us.dataAccess.weibo.GetListByFollower(sourceNode.ID)
	if err != nil {
		return nil, nil, err
	}

	times--

OutLoop:
	for i, record := range records {
		fmt.Printf("Times: %d, Record %d \n", times, i)
		if i >= 10 {
			return nodes, links, nil
		}

		if record.ID == mainNode.ID {
			sourceNode.Category = 3
			continue OutLoop
		}

		for _, node := range nodes {
			if record.ID == node.ID {
				links = append(links, &Link{Source: record.NickName, Target: sourceNode.Name, Weight: (times + 1) * 3})
				continue OutLoop
			}
		}

		recordNode := &Node{ID: record.ID, Category: 1, Name: record.NickName, Value: (times + 1) * 3}
		nodes = append(nodes, recordNode)
		links = append(links, &Link{Source: record.NickName, Target: sourceNode.Name, Weight: (times + 1) * 3})

		if times > 0 {
			nodes, links, err = us.getRelationCircle(sourceNode, recordNode, nodes, links, times)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return nodes, links, nil
}

func (us *Service) getZhihuRelationCircle(mainNode *Node, sourceNode *Node, nodes []*Node, links []*Link, times int) ([]*Node, []*Link, error) {
	records, err := us.dataAccess.zhihu.GetListByFollower(sourceNode.ZhihuID)
	if err != nil {
		return nil, nil, err
	}

	times--

OutLoop:
	for i, record := range records {
		if i >= 10 {
			return nodes, links, nil
		}

		if record.ID == mainNode.ID {
			sourceNode.Category = 3
			continue OutLoop
		}

		for _, node := range nodes {
			if record.ID == node.ID {
				links = append(links, &Link{Source: record.NickName, Target: sourceNode.Name, Weight: (times + 1) * 3})
				continue OutLoop
			}
		}

		recordNode := &Node{ID: record.ID, Category: 1, Name: record.NickName, Value: (times + 1) * 3}
		nodes = append(nodes, recordNode)
		links = append(links, &Link{Source: record.NickName, Target: sourceNode.Name, Weight: (times + 1) * 3})

		if times >= 0 {
			nodes, links, err = us.getRelationCircle(sourceNode, recordNode, nodes, links, times)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return nodes, links, nil
}
