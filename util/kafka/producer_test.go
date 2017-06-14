package kafka

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	assert.NotPanics(t, func() { GetProducer() })

	assert.NoError(t, GetProducer().SendMessage("bigdata", "weibo-52436246-nickname"))
	GetProducer().SendMessage("bigdata", "zhihu-52141325246-nickname")
	GetProducer().SendMessage("bigdata", "zhihu-521231225246-nickname")
}
