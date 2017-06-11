package model

import (
	"testing"

	"brush/core/db"

	"github.com/stretchr/testify/assert"
)

func TestZhihuRelation(t *testing.T) {
	db.GetZhihuInstance().OpenConnect()

	assert.NotPanics(t, func() { ZhihuRelationDataIns = NewZhihuRelationData() })
}
