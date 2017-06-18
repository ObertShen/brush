package model

import (
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert := assert.New(t)
	assert.NotPanics(func() { UserDataIns = NewUserData() })

	// testRecord := &User{
	// 	Name:      "123",
	// 	NickName:  "xiang",
	// 	UnionID:   util.GetUUID(),
	// 	Role:      "admin",
	// 	Password:  util.GetSHA256String("123456"),
	// 	UpdatedAt: time.Now().Unix(),
	// 	CreatedAt: time.Now().Unix(),
	// }

	// affected, err := UserDataIns.Insert(testRecord)
	// if assert.NoError(err) {
	// 	assert.NotZero(affected)
	// }

	has, err := UserDataIns.Get(&User{Name: "lpx"})
	if assert.NoError(err) {
		assert.True(has)
	}

	userList, err := UserDataIns.GetList(&User{ID: 1})
	if assert.NoError(err) {
		assert.True(len(userList) == 1)
	}
}
