package usecases

import (
	"example.com/gin_wire/app/repositories"
)

type SampleUsecase struct {
	sampleRepository *repositories.SampleRepository
}

func NewSampleUsecase(
	sampleRepository *repositories.SampleRepository,
) *SampleUsecase {
	return &SampleUsecase{
		sampleRepository: sampleRepository,
	}
}

func (uc *SampleUsecase) Get() string {
	return uc.sampleRepository.Get()
}
