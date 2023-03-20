package data

import (
	"go-server/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var DataSet = wire.NewSet(NewTrackerRepo ,InitDatabase)

// Data .
type Data struct {
	// TODO wrapped database client
	g *gorm.DB
	log.Logger
}

// NewData .
func NewData(c *conf.Data,g*gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{g :InitDatabase(c)}, cleanup, nil
}
