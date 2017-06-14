package model

import (
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"brush/core/db"
	"brush/util"

	"github.com/stretchr/testify/assert"
)

func TestWeibo(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert.NotPanics(t, func() { WeiboDataIns = NewWeiboData() })
	util.ReadLine("./weibo.txt", addWeibo)
}

func addWeibo(record string) {
	columns := strings.Split(record, "\t\t")
	weiboID, _ := strconv.ParseInt(columns[0], 10, 0)
	userID, _ := strconv.ParseInt(columns[1], 10, 0)
	commentNum, _ := strconv.ParseInt(columns[2], 10, 0)
	likeNum, _ := strconv.ParseInt(columns[3], 10, 0)
	sentTime, _ := strconv.ParseInt(columns[4], 10, 0)
	content := ""

	if columns[5] != "" {
		for _, value := range columns[5] {
			_, size := utf8.DecodeRuneInString(string(value))
			if size <= 3 {
				content += string(value)
			}
		}
	}

	WeiboDataIns.Insert(&Weibo{
		ID:            weiboID,
		UserID:        userID,
		CommentNumber: commentNum,
		LikeNumber:    likeNum,
		SentTime:      sentTime,
		Content:       content,
		CreatedAt:     time.Now().Unix(),
		UpdatedAt:     time.Now().Unix(),
	})
}
