package user

import (
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestGetWeiboUserTags(t *testing.T) {
	db.GetInstance().OpenConnect()

	assert.NotPanics(t, func() { ServiceIns = GetServiceIns() })

	result, err := ServiceIns.GetWeiboUserTags(10)
	if assert.NoError(t, err) {
		assert.NotNil(t, result)
	}

	// result, err = ServiceIns.GetZhihuUserTags("mileijun")
	// if assert.NoError(t, err) {
	// 	assert.NotNil(t, result)
	// }
}
