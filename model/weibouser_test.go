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

func TestWeiboUser(t *testing.T) {
	db.GetInstance().OpenConnect()

	assert.NotPanics(t, func() { WeiboUserDataIns = NewWeiboUserData() })

	util.ReadLine("./info.txt", addWeiboUser)
}

func addWeiboUser(record string) {
	columns := strings.Split(record, "\t\t")
	userID, _ := strconv.ParseInt(columns[0], 10, 0)
	followNumber, _ := strconv.ParseInt(columns[1], 10, 0)
	fansNumber, _ := strconv.ParseInt(columns[2], 10, 0)
	weiboNumber, _ := strconv.ParseInt(columns[3], 10, 0)
	introduction := ""

	if columns[8] != "" {
		for _, value := range columns[8] {
			_, size := utf8.DecodeRuneInString(string(value))
			if size <= 3 {
				introduction += string(value)
			}
		}
	}

	gender := 0
	if columns[6] == "男" {
		gender = 1
	} else if columns[6] == "女" {
		gender = 2
	}

	WeiboUserDataIns.Insert(&WeiboUser{
		ID:           userID,
		FollowNumber: followNumber,
		FansNumber:   fansNumber,
		WeiboNumber:  weiboNumber,
		NickName:     columns[4],
		Region:       columns[5],
		Gender:       gender,
		Birthday:     columns[7],
		Introduction: introduction,
		RegisterTime: columns[9],
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
	})

	// if err != nil {
	// 	log.Panic(err)
	// }
}
