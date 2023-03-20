package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TrackerRepo struct {
	db *gorm.DB
	log  *log.Helper
}

type Block struct{

}

// NewGreeterRepo .
func NewTrackerRepo(g *gorm.DB, logger log.Logger)	*TrackerRepo {
	return &TrackerRepo{
		db: g,
		log:  log.NewHelper(logger),
	}
}

func (t *TrackerRepo) CreateBlockDB() (*Block,*gorm.DB){
	return nil,nil;
}

