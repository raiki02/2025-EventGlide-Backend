package ioc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/raiki02/EG/internal/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitDB() *gorm.DB {
	dsn := viper.GetString("mysql.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqldb.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	sqldb.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

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
		&model.Post{},
		&model.PostDraft{},
		&model.Feed{},
		&model.Approvement{},
		&model.AuditorForm{},
	)
}
