package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DBConn provides the general connection to the database
	DBConn *gorm.DB
)
