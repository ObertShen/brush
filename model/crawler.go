package model

import "log"

// Crawler 与数据库中的 crawler 表相对应 用于存储项目的巡查信息
type Crawler struct {
	ID        int64 `xorm:"Int(11) pk autoincr" json:"id"`
	UserID    int64 `xorm:"Int(11) notnull" json:"-"`
	IsAdmin   bool  `xorm:"notnull" json:"isAdmin"`
	UpdatedAt int64 `xorm:"Int(11) notnull" json:"updatedAt"`
	CreatedAt int64 `xorm:"Int(11) notnull" json:"createdAt"`
	IsDeleted bool  `xorm:"notnull" json:"-"`
}

var (
	// CrawlerDataIns 是 CrawlerData 类的单例
	CrawlerDataIns *CrawlerData
)

// CrawlerData is a class that includeing fundamental functions for CURD operations in crawler table
type CrawlerData struct {
	*DBConn
}

// GetCrawlerDataIns 用于获取 CrawlerData 类的单例
func GetCrawlerDataIns() *CrawlerData {
	if CrawlerDataIns == nil {
		CrawlerDataIns = NewCrawlerData()
	}

	return CrawlerDataIns
}

// NewCrawlerData init CrawlerData class
func NewCrawlerData() *CrawlerData {
	cd := &CrawlerData{GetConnIns()}
	if err := cd.conn.Sync2(new(Crawler)); err != nil {
		log.Panic(err)
	}
	return cd
}

// GetList 从 crawler 中查询纪录
func (cd *CrawlerData) GetList(params *Crawler) (crawlerList []*Crawler, err error) {
	crawlerList = []*Crawler{}
	if err = cd.conn.UseBool("is_deleted", "is_admin").Find(&crawlerList, params); err != nil {
		return
	}

	return
}

// Get 从 crawler 中查询纪录
func (cd *CrawlerData) Get(params *Crawler) (bool, error) {
	return cd.conn.UseBool("is_deleted", "is_admin").Get(params)
}

// Insert 从 crawler 中增加纪录
func (cd *CrawlerData) Insert(crawler *Crawler) (int64, error) {
	return cd.conn.Insert(crawler)
}
