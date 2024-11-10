package service

import "main/model/manage"

func (s *Service) GetItemsReplacementService(data *[]manage.ReplacementItem) error {
	return s.Repo.GetItemsReplacementRepo(data)
}
