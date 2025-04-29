package repository

import (
	"context"
	"xws/webook/internal/domain"
	"xws/webook/internal/repository/dao"
)


type UserRepository struct {
	dao *dao.UserDAO
}

func NewRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}


func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

