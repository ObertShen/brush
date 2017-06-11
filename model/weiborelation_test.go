package model

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"brush/core/db"
	"brush/util"

	"github.com/stretchr/testify/assert"
)

func TestWeiboRelation(t *testing.T) {
	db.GetDefaultInstance().OpenConnect()

	assert.NotPanics(t, func() { WeiboRelationDataIns = NewWeiboRelationData() })

	util.ReadLine("./relation.txt", addWeiboRelation)
}

func addWeiboRelation(record string) {
	fmt.Println("========")
	fmt.Println(record)

	columns := strings.Split(record, "\t\t")
	fmt.Println(columns[0])
	fmt.Println(columns[1])

	userID, _ := strconv.ParseInt(columns[0], 10, 0)
	followID, _ := strconv.ParseInt(columns[1], 10, 0)

	fmt.Println(len(columns))
	fmt.Println(userID)
	fmt.Println(followID)
}
