package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
)

type postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var db *gorm.DB

func init() {

	file, _ := ioutil.ReadFile("./conf/db.json")

	pg := postgres{}

	_ = json.Unmarshal(file, &pg)

	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pg.User, pg.Password, pg.Host, pg.Dbname)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Product{})

}

func getDB() *gorm.DB {
	return db
}
