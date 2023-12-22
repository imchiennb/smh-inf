package usecase

import (
	"github.com/imchiennb/acmex/internal/app/model"
	"gorm.io/gorm"
)

type EntityUsecase struct {
}

type EntityRepository interface {
	CreateEntity(entity *model.Entity) (error, model.Entity)
	UpdateEntity(entity *model.Entity) (error, model.Entity)
	ReadEntity(entity *model.Entity) (error, model.Entity)
	DeleteEntity(entity *model.Entity) (error, model.Entity)
	ListEntities() (error, []model.Entity)
	PageEntities(limit int, offset int) (error, []model.Entity)
}

func NewEntityUsecase(db *gorm.DB) *EntityUsecase {
	return &EntityUsecase{}
}
