package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/ayotask1?charset=utf8mb4&parseTime=True&loc=Local", "ramana", "MvRamana@2022")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
