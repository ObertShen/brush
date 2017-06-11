package model

import (
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestWeibo(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert.NotPanics(t, func() { WeiboDataIns = NewWeiboData() })
}
