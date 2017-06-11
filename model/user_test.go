package model

import (
	"testing"
	"time"

	"brush/core/db"
	"brush/util"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert := assert.New(t)
	assert.NotPanics(func() { UserDataIns = NewUserData() })

	testRecord := &User{
		Name:      "lpx",
		NickName:  "xiang",
		UnionID:   util.GetUUID(),
		Role:      "admin",
		Password:  util.GetSHA256String("123456"),
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}

	affected, err := UserDataIns.Insert(testRecord)
	if assert.NoError(err) {
		assert.NotZero(affected)
	}

	has, err := UserDataIns.Get(&User{ID: testRecord.ID})
	if assert.NoError(err) {
		assert.True(has)
	}

	userList, err := UserDataIns.GetList(&User{ID: testRecord.ID})
	if assert.NoError(err) {
		assert.True(len(userList) == 1)
	}
}
