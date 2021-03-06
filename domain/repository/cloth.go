package repository

import "github.com/kons16/team7-backend/domain/entity"

type Cloth interface {
	Create(cloth *entity.Cloth) (int, error)
	GetAll() *[]entity.Cloth
	CreateUserCloth(userID int, clothID int) error
	GetBuyCloth(userID int) *[]entity.Cloth
	GetByIDs(clothIDs []int) *[]entity.Cloth
}
