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
	// Zhihu 知乎信息
	Zhihu struct {
		model.ZhihuUser
		Type string `json:"type"`
	}
	// Weibo 知乎信息
	Weibo struct {
		model.WeiboUser
		Type string `json:"type"`
	}
	// Node 构造关系图节点
	Node struct {
		Category int    `json:"category"`
		Name     string `json:"name"`
		Value    int    `json:"value"`
		Label    string `json:"label,omitempty"`
	}
	// Link 构造关系图的连接
	Link struct {
		Source string `json:"source"`
		Target string `json:"target"`
		Weight int    `json:"weight"`
		Name   string `json:"name,omitempty"`
	}
)
