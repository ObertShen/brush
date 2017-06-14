package model

import (
	"fmt"
	"testing"
	"time"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestUserTag(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert := assert.New(t)
	assert.NotPanics(func() { userTagDataIns = NewUserTagData() })

	testRecord := &UserTag{
		Platform:  "weibo",
		UserID:    10,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	testRecord.Tags = append([]nameAndValue{}, nameAndValue{Name: "test", Value: 22})

	affected, err := userTagDataIns.Insert(testRecord)
	if assert.NoError(err) {
		assert.NotZero(affected)
	}

	has, err := userTagDataIns.Get(&UserTag{ID: testRecord.ID})
	if assert.NoError(err) {
		assert.True(has)
	}

	userList, err := userTagDataIns.GetList(&UserTag{ID: testRecord.ID})
	if assert.NoError(err) {
		assert.True(len(userList) == 1)
	}

	tags, err := userTagDataIns.GetWeiboTags(10, "test")
	assert.NoError(err)
	fmt.Println(tags)
}
