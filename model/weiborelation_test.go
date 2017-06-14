package model

import (
	"strconv"
	"strings"
	"testing"

	"brush/core/db"
	"brush/util"

	"github.com/stretchr/testify/assert"
)

func TestWeiboRelation(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert.NotPanics(t, func() { WeiboRelationDataIns = NewWeiboRelationData() })

	util.ReadLine("./relation.txt", addWeiboRelation)
}

func addWeiboRelation(record string) {
	columns := strings.Split(record, "\t\t")

	s := strings.Split(columns[1], "\n")
	columns[1] = s[0]
	userID, _ := strconv.ParseInt(columns[0], 10, 0)
	followID, _ := strconv.ParseInt(columns[1], 10, 0)

	WeiboRelationDataIns.Insert(&WeiboRelation{
		FollowID:   userID,
		FollowerID: followID,
	})
}
