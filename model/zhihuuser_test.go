package model

import (
	"fmt"
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestZhihuUser(t *testing.T) {
	db.GetInstance().OpenConnect()

	assert.NotPanics(t, func() { ZhihuUserDataIns = NewZhihuUserData() })

	userList, err := ZhihuUserDataIns.GetListByFollower("wang-ni-ma-94")
	if assert.NoError(t, err) {
		fmt.Println(len(userList))
		assert.NotZero(t, len(userList))
	}
}
