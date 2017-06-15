package main

import (
	"runtime"

	brush "brush/controller"
	mysql "brush/core/db"
	"brush/model"
	// "brush/util/kafka"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// kafka.GetProducer()
	weiboDB := mysql.GetDefaultInstance()
	weiboDB.OpenConnect()
	defer weiboDB.CloseConnect()

	dbcon := mysql.GetZhihuInstance()
	dbcon.OpenConnect()
	defer dbcon.CloseConnect()

	model.InitDB()
}

func main() {
	brush.Mount().Run(":3000")
}
