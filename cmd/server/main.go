package main

import (
	"github.com/shubhamrwtktw/lic/cmd/app"
	"github.com/shubhamrwtktw/lic/internal/config/db"
	"github.com/shubhamrwtktw/lic/internal/config/env"
)


func main (){
// load config
	env.LoadEnv()

	// connect db

	db.Connect()

	// init app

	app.InitApp()


}