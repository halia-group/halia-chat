package dao

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"halia-chat/server/model"
)

type Dao interface {
	Register(ctx context.Context, username, password, nickname string) error
	Login(ctx context.Context, username, password string) (*model.User, error)
}

type dao struct {
	db  *gorm.DB
	log *log.Entry
}

func New() (Dao, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	logger := log.WithField("component", "dao")
	logger.Infoln("initialized")
	return &dao{db: db, log: logger}, nil
}
