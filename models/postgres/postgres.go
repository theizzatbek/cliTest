package postgres

import (
	"cliTest/config"
	"cliTest/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {

	pg := config.GetInstance().Postgres

	dbUri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pg.User, pg.Password, pg.Host, pg.Dbname)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.User{}, &models.Product{})

}

func GetDB() *gorm.DB {
	return db
}
