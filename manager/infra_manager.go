package manager

import (
	"log"

	"ticket.narindo.com/config"
	"ticket.narindo.com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func InitInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserRole{},
		&model.Priority{},
		&model.Status{},
		&model.PicDepartment{},
		&model.Pic{},
		&model.Ticket{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
