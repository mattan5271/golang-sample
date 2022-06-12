package models

import (
	"fmt"
	"golang-sample/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_ad"`
	UpdatedAt time.Time  `json:"updated_ad"`
	DeletedAt *time.Time `json:"deleted_ad"`
}

type User struct {
	Model        // id, created_at, updated_at, deleted_atを持ってくれる
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Books []Book `gorm:"foreignKey:UserID" json:"books"`
}

type Book struct {
	Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

var Db *gorm.DB

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)
	fmt.Println(dbConnectInfo)

	// configから読み込んだ情報を元に、データベースに接続します
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Successfully connect database..")
	}

	// 接続したデータベースにusersテーブルを作成
	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{}, &Book{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table..")
	}
}
