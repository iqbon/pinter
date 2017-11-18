package main

import (
	_ "pinter/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"pinter/models"
	"pinter/models/database"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "host=localhost port=5432 user=postgres dbname=pinter sslmode=disable")
}

func main() {
	database.DB.AutoMigrate(&models.Soal{}, models.Jawaban{})
	defer database.DB.Close()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
