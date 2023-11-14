package usecase

import (
	"fmt"

	model "github.com/DavidAfdal/Weekly-FinderJob-Be/models"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/models/request"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/models/response"
	repository "github.com/DavidAfdal/Weekly-FinderJob-Be/repositorys"
)

type JobUseCase struct {
	JobRepo *repository.JobRepository
}

func NewJobsUseCase(repo *repository.JobRepository) *JobUseCase {
	return &JobUseCase{JobRepo: repo}
}

func (u *JobUseCase) Create(job request.CreateJobInput) {

	jobModel := model.Job{
		Title:        job.Title,
		Description:  job.Description,
		Company:      job.Company,
		ImageCompany: job.ImageCompany,
		Category:     job.Category,
		Status:       job.Status,
		Location:     job.Location,
		Salary:       job.Salary,
		UserId:       job.UserId,
	}

	u.JobRepo.Save(jobModel)
}

func (u *JobUseCase) FindAll() []response.JobResponse {

	result := u.JobRepo.FindAll()
	fmt.Println(result)

	var jobs []response.JobResponse

	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else {
		for _, value := range result {
			job := response.JobResponse{
				Id:           value.Id,
				Title:        value.Title,
				Company:      value.Company,
				ImageCompany: value.ImageCompany,
				Status:       value.Status,
				Category:     value.Category,
				CreatedAt:    value.CreatedAt,
			}
			jobs = append(jobs, job)
		}
	}

	return jobs
}

func ( u *JobUseCase) FindById(id int) (model.Job, error) {
	result, err := u.JobRepo.FindById(id)

	return result, err
}

func (u *JobUseCase) FindByUserId(userId string) []response.JobResponse {

	result := u.JobRepo.FindByUserId(userId)
	fmt.Println(result)

	var jobs []response.JobResponse
	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else {
		for _, value := range result {
			job := response.JobResponse{
				Id:           value.Id,
				Title:        value.Title,
				Company:      value.Company,
				ImageCompany: value.ImageCompany,
				Status:       value.Status,
				Category:     value.Category,
				CreatedAt:    value.CreatedAt,
			}
			jobs = append(jobs, job)
		}
	}

	return jobs
}

func (u *JobUseCase) FindByCategory(category string) []response.JobResponse {

	result := u.JobRepo.FindByCategory(category)
	var jobs []response.JobResponse
	if len(result) == 0 {
		jobs = []response.JobResponse{}
	} else {
		for _, value := range result {
			job := response.JobResponse{
				Id:           value.Id,
				Title:        value.Title,
				Company:      value.Company,
				ImageCompany: value.ImageCompany,
				Status:       value.Status,
				Category:     value.Category,
				CreatedAt:    value.CreatedAt,
			}
			jobs = append(jobs, job)
		}
	}

	return jobs
}

func (u *JobUseCase) Update(job request.UpdateJobInput) error {
	jobData, err := u.JobRepo.FindById(job.Id)

	if err != nil {
		return err
	}

	jobData.Title = job.Title
	jobData.Company = job.Company
	jobData.Description = job.Description
	jobData.ImageCompany = job.ImageCompany
	jobData.Category = job.Category
	jobData.Status = job.Status
	jobData.Location = job.Location
	jobData.Salary = job.Salary

	u.JobRepo.Update(jobData)

	return nil
}

func (u *JobUseCase) Delete(jobId int) error {

	_, err := u.JobRepo.FindById(jobId)

	if err != nil {
		return err
	}

	u.JobRepo.Delete(jobId)

	return nil
}