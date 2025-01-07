package initialize

import (
	"fmt"
	"go-services/global"
	"go-services/internal/po"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresql() {
	p := global.Config.Postgresql
	dsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	s := fmt.Sprintf(dsn, p.Host, p.User, p.Password, p.DBName, p.Port)
	db, err := gorm.Open(postgres.Open(s), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	global.PDB = db

	// Auto migrate
	migrateTables()
}

func migrateTables() {
	err := global.PDB.AutoMigrate(&po.User{}, &po.Task{})
	if err != nil {
		panic("failed to migrate database")
	}
}
