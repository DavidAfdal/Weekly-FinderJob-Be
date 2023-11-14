package repository

import (
	"errors"

	helper "github.com/DavidAfdal/Weekly-FinderJob-Be/helpers"
	model "github.com/DavidAfdal/Weekly-FinderJob-Be/models"
	"gorm.io/gorm"
)

type ApplierRepository struct {
	DB *gorm.DB
}


func NewApplierRepository(db *gorm.DB) *ApplierRepository {
	return &ApplierRepository{DB: db}
}


func (r *ApplierRepository) Save(applier model.Applier){
	result := r.DB.Create(&applier)
	helper.ErrorPanic(result.Error)
}

func (r *ApplierRepository) FindByUserId(userId string) (model.Applier, error) {
	var applier model.Applier
	result := r.DB.Where("user_id = ?", userId).Preload("Jobs").First(&applier)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return applier, errors.New("Applier Not Found")
	} else {
		return applier, nil
	}

}

func (r *ApplierRepository) AppendJob(job model.Job) {
	r.DB.Model(&model.Applier{}).Association("Jobs").Append(&job)
}
