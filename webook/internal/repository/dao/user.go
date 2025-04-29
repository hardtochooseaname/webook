package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (ud *UserDAO) Insert(ctx context.Context, u User) error {
	u.Ctime = time.Now().UnixMilli()
	u.Utime = u.Ctime
	return ud.db.WithContext(ctx).Create(&u).Error
}

// 连通 GORM，直接对应数据库中的 user 表
// 这个映射到数据库中的数据结构在其他地方也有别的叫法，例如 entity、model、PO（持久化对象）等
type User struct {
	ID       int64 `gorm:"primaryke;autoIncrement"`
	Email    string `gorm:"unique;not null"`
	Password string

	Ctime int64
	Utime int64
}
