package server

import (
	"context"
	"errors"
	"sync"
)

type Dao interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) error
}

type dao struct {
	users map[string]string
	lock  sync.RWMutex
}

func newDao() *dao {
	return &dao{
		users: make(map[string]string),
	}
}

func (d *dao) Register(ctx context.Context, username, password string) error {
	d.lock.RLock()
	_, exists := d.users[username]
	d.lock.RUnlock()
	if exists {
		return errors.New("账号已存在")
	}
	d.lock.Lock()
	d.users[username] = password
	d.lock.Unlock()
	return nil
}

func (d *dao) Login(ctx context.Context, username, password string) error {
	if username == "" || password == "" {
		return errors.New("账号或密码不能为空")
	}
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.users[username] != password {
		return errors.New("账号或密码错误")
	}
	return nil
}
