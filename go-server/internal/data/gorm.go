package data

import (
	"go-server/internal/conf"
	"log"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(cfg *conf.Data)*gorm.DB{
	var glog = logger.New(log.New(log.Writer(),"\r\n",log.Flags()),logger.Config{
		SlowThreshold: time.Second,
		LogLevel: logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful: true,
	})
	db,err := gorm.Open(mysql.New(mysql.Config{
		DSN:cfg.Database.Source+"?parseTime=true",
	}),&gorm.Config{
		Logger: glog,
	})
	if err !=nil{
		log.Panic(err)
		panic("unabel to connect to database );")
	}
	return db
}