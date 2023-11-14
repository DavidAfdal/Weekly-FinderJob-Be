package repository

import (
	"errors"
	"fmt"

	helper "github.com/DavidAfdal/Weekly-FinderJob-Be/helpers"
	model "github.com/DavidAfdal/Weekly-FinderJob-Be/models"
	"gorm.io/gorm"
)

type JobRepository struct {
	Db *gorm.DB
}

func NewJobRepository(Db *gorm.DB) *JobRepository {
	return &JobRepository{Db: Db}
}

func (r *JobRepository) Save(job model.Job) {
	result := r.Db.Create(&job)
	helper.ErrorPanic(result.Error)
}

func (r *JobRepository) FindAll() []model.Job {
	var jobs []model.Job

	result := r.Db.Find(&jobs)
	helper.ErrorPanic(result.Error)

	return jobs
}

func (r *JobRepository) FindById(jobId int) (model.Job, error) {
	var job model.Job

	result := r.Db.First(&job, jobId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return job, errors.New("Job Not Found")
	} else {
		return job, nil
	}
}

func (r *JobRepository) FindByUserId(userId string) []model.Job {
	var jobs []model.Job

	result := r.Db.Where("user_id = ?", userId).Find(&jobs)
	fmt.Println(result)
	helper.ErrorPanic(result.Error)

	return jobs
}

func (r *JobRepository) FindByCategory(category string) []model.Job {
	var jobs []model.Job

	result := r.Db.Where("category = ?", category).Find(&jobs)
	helper.ErrorPanic(result.Error)

	return jobs
}

func (r *JobRepository) Update(job model.Job) {
	result := r.Db.Model(&model.Job{}).Where("id = ?", job.Id).Updates(job)
	helper.ErrorPanic(result.Error)
}

func (r *JobRepository) Delete(jobId int) {
	var job model.Job
	result := r.Db.Where("id = ?", jobId).Delete(&job)
	helper.ErrorPanic(result.Error)
}
