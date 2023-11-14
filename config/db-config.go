package config

import (
	helper "github.com/DavidAfdal/Weekly-FinderJob-Be/helpers"
	model "github.com/DavidAfdal/Weekly-FinderJob-Be/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/jobfinder_db?charset=utf8&parseTime=True&loc=Local"))
    helper.ErrorPanic(err)

	db.Migrator().DropTable("jobs")
	db.Migrator().DropTable("appliers")
	db.Migrator().DropTable("applier_job")

	db.AutoMigrate(&model.Applier{})
	db.AutoMigrate(&model.Job{})
	return db
}