package usecase

import (
	"errors"

	model "github.com/DavidAfdal/Weekly-FinderJob-Be/models"
	"github.com/DavidAfdal/Weekly-FinderJob-Be/models/request"
	repository "github.com/DavidAfdal/Weekly-FinderJob-Be/repositorys"
)

type ApplierUseCase struct {
	ApplierRepo *repository.ApplierRepository
	JobRepo     *repository.JobRepository
}

func NewApplierUseCase(ApplierRepo *repository.ApplierRepository, jobRepo *repository.JobRepository) *ApplierUseCase {
	return &ApplierUseCase{ApplierRepo: ApplierRepo, JobRepo: jobRepo}
}

func (u *ApplierUseCase) ApplyJob(applier request.ApplierRequest) error {

	jobData, err := u.JobRepo.FindById(applier.JobId)

	if err != nil {
		return err
	}

	if applier.UserId == jobData.UserId {
		return errors.New("Cannot Apply")
	}

	if _, err := u.ApplierRepo.FindByUserId(applier.UserId); err != nil {

		applierData := model.Applier{
			Name:   applier.Name,
			UserID: applier.UserId,
			Jobs:   []model.Job{jobData},
		}
		u.ApplierRepo.Save(applierData)
	} else {

		u.ApplierRepo.AppendJob(jobData)
	}

return nil;
}

func (u *ApplierUseCase) FindByUserId(userId string) (model.Applier, error) {
	data, err := u.ApplierRepo.FindByUserId(userId)
	return data, err
}