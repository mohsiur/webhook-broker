package storage

import (
	"github.com/imyousuf/webhook-broker/config"
	"github.com/imyousuf/webhook-broker/storage/data"
)

// DataAccessor is the facade to all the data repository
type DataAccessor interface {
	GetAppRepository() AppRepository
	Close()
}

// AppRepository allows storage operation interaction for App
type AppRepository interface {
	GetApp() (*data.App, error)
	StartAppInit(data *config.SeedData) error
	CompleteAppInit() error
}

// ProducerRepository allows storage operation interaction for Producer
type ProducerRepository interface {
	Store(producer *data.Producer) (*data.Producer, error)
	Get(producerID string) (*data.Producer, error)
	GetList(page *data.Pagination) ([]*data.Producer, *data.Pagination, error)
}
