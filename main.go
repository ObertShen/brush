package main

import (
	"runtime"

	brush "brush/controller"
	mysql "brush/core/db"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	weiboDB := mysql.GetDefaultInstance()
	weiboDB.OpenConnect()
	defer weiboDB.CloseConnect()

	dbcon := mysql.GetZhihuInstance()
	dbcon.OpenConnect()
	defer dbcon.CloseConnect()

	brush.Mount().Run(":3000")
}
