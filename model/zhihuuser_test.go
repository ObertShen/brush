package model

import (
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestZhihuUser(t *testing.T) {
	db.GetInstance().OpenConnect()

	assert.NotPanics(t, func() { ZhihuUserDataIns = NewZhihuUserData() })
}
