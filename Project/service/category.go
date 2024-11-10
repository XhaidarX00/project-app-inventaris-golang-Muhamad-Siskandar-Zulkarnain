package service

import (
	"main/model/category"
)

func (s *Service) GetCategoriesService(category *[]category.Category) error {
	err := s.Repo.GetCategoriesRepo(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddCategoryService(category *category.Category) error {
	err := s.Repo.AddCategoryRepo(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCategoryByIdService(category *category.Category) error {
	err := s.Repo.GetCategoryByIdRepo(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PutCategoryByIdService(category *category.Category) error {
	err := s.Repo.PutCategoryByIdRepo(category)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeletCategoryByIdService(id int) error {
	err := s.Repo.DeletCategoryByIdRepo(id)
	if err != nil {
		return err
	}
	return nil
}
