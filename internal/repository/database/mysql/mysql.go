package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Velocyes/mini-go-project/internal/consts"
	"github.com/Velocyes/mini-go-project/internal/model"
)

type mySql struct {
	dbInstance *gorm.DB
}

func InitMySQL(cfg *model.Config) (mySqlInstance *mySql, err error) {
	if cfg == nil {
		return nil, consts.ErrNilConfig
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Product{}, &model.Order{}, &model.OrderDetail{})

	return &mySql{
		dbInstance: db,
	}, nil
}
