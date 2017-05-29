package main

import (
	"runtime"

	brush "brush/controller"
	mysql "brush/core/db"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dbcon := mysql.GetInstance()
	dbcon.OpenConnect()
	defer dbcon.CloseConnect()

	brush.Mount().Run(":3000")
}
