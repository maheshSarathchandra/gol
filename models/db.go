package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //PostgreSQL Driver
)

var ormObject orm.Ormer

func ConnectToDb() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=AdminUserData dbname=postgres host=localhost sslmode=disable")
	orm.RegisterModel(new(Users))
	ormObject = orm.NewOrm()
}

func GetOrmObject() orm.Ormer {
	return ormObject
}
