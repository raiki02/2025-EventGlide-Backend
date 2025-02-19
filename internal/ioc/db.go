package ioc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/raiki02/EG/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	sqldb.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	sqldb.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))

	err = migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Activity{},
		&model.ActivityDraft{},
		&model.Comment{},
		&model.SubComment{},
		&model.Post{},
		&model.PostDraft{},
		&model.Number{},
	)
}
