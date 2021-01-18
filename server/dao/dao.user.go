package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"halia-chat/server/model"
	"time"
)

func (d dao) Register(ctx context.Context, username, password, nickname string) error {
	if len(username) == 0 || len(password) == 0 || len(nickname) == 0 {
		return errors.New("参数不能为空")
	}
	var user model.User
	err := d.db.Where("username=?", username).Take(&user).Error
	if err == nil { // 账号存在
		return errors.New("账号已存在")
	}
	if err != gorm.ErrRecordNotFound {
		d.log.WithError(err).Warnln("query failed")
		return errors.New("服务器内部错误")
	}
	// 开始注册
	user.Username = username
	user.Password = password
	user.Nickname = nickname
	user.CreatedAt = time.Now()
	if err := d.db.Create(&user).Error; err != nil {
		d.log.WithError(err).Warnln("create failed")
		return errors.New("注册失败")
	}
	return nil
}

func (d dao) Login(ctx context.Context, username, password string) (*model.User, error) {
	var user model.User
	err := d.db.Where("username=?", username).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("账号或密码错误")
	}
	if err != nil {
		d.log.WithError(err).Warnln("query failed")
		return nil, errors.New("登录失败")
	}
	if user.Password != password {
		return nil, errors.New("账号或密码错误")
	}
	return &user, nil
}
