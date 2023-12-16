package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user_microservices/common"
)

type PostgresDBWaizly struct {
	DB *gorm.DB
}

func (dbS *PostgresDBWaizly) Init() error {

	db, err := gorm.Open("postgres", "host="+common.Config.DbAddrsWaizlyPublic+" port="+common.Config.DBPortWaizly+" user="+common.Config.DbUsernameWaizly+" dbname="+common.Config.DbNameWaizly+" sslmode=disable password="+common.Config.DbPasswordWaizly+"")
	//fmt.Println(common.Config.DbAddrsWarnaCRMPublic)
	db.DB().Ping()
	db.DB().SetMaxIdleConns(0)

	if err != nil {
		return err
	}

	dbS.DB = db

	return err
}
