package service

import (
	"main/model/investment"
)

func (s *Service) GetTotalInvesmentService(invesment *investment.Investment) error {
	err := s.Repo.GetTotalInvesmentsRepo(invesment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTotalInvesmentByIdService(invesment *investment.ItemInvestment) error {
	err := s.Repo.GetTotalInvesmentsByIdRepo(invesment)
	if err != nil {
		return err
	}
	return nil
}
