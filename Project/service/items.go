package service

import (
	"main/model/manage"
	"main/model/response"
)

func (r *Service) GetItemsPaginatedService(data *response.PaginationResponse) error {
	return r.Repo.GetItemsPaginatedRepo(data)
}

func (s *Service) AddInventoryItemService(item *manage.Item) error {
	return s.Repo.AddInventoryItemRepo(item)
}

func (s *Service) GetInventoryItemByIdService(item *manage.Item) error {
	return s.Repo.GetInventoryItemsByIDRepo(item)
}

func (s *Service) UpdateInventoryItemByIdService(item *manage.Item) error {
	return s.Repo.UpdateInventoryItemByIdRepo(item)
}

func (s *Service) DeleteInventoryItemByIdService(id int) error {
	return s.Repo.DeleteInventoryItemByIdRepo(id)
}
