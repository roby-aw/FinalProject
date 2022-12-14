package service

import (
	"final-project/config"
	"final-project/dto"
	"final-project/interfaces"
)

type SosmedService struct {
	SosmedRepository interfaces.SosmedRepository
	conf             config.AppConfig
}

func NewSosmedService(sosmedRepository interfaces.SosmedRepository, conf config.AppConfig) interfaces.SosmedService {
	return &SosmedService{
		SosmedRepository: sosmedRepository,
		conf:             conf,
	}
}

func (s *SosmedService) CreateSosmed(sosmed *dto.CreateSosmed) (dto.CreateSosmed, error) {
	sos, err := s.SosmedRepository.CreateSosmed(sosmed)

	if err != nil {
		return sos, err
	}

	return sos, nil
}

func (s *SosmedService) GetSosmed() ([]dto.GetAllSosmed, error) {
	return s.SosmedRepository.GetSosmed()
}

func (s *SosmedService) UpdateSosmed(id uint, sosmed *dto.UpdateSosmed) (dto.UpdateSosmed, error) {
	return s.SosmedRepository.UpdateSosmed(id, sosmed)
}

func (s *SosmedService) DeleteSosmed(id uint) error {
	return s.SosmedRepository.DeleteSosmed(id)
}

func (s *SosmedService) GetSosmedByID(id uint) (dto.GetAllSosmed, error) {
	return s.SosmedRepository.GetSosmedByID(id)
}
