package service

import (
	"main/library"
	users "main/model/user"
	"net/http"
	"time"
)

func (s *Service) RegisterService(user users.User) error {
	err := s.Repo.RegisterRepo(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginService(user *users.User) error {
	err := s.Repo.LoginRepo(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) TokenCheck(token string) string {
	err := s.Repo.TokenCheckRepo(token)
	if err != "" {
		return err
	}

	return ""
}

// Fungsi untuk membersihkan token yang sudah kadaluarsa
func (s *Service) CleanExpiredTokens(w http.ResponseWriter) bool {
	err := s.Repo.CleanExpiredTokensRepo()
	if err != "" {
		response := library.UnauthorizedRequest()
		library.JsonResponse(w, response)
		return false
	}

	return true
}

func (s *Service) CheckToken() {
	ticker := time.NewTicker(12 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := s.Repo.CleanExpiredTokensRepo(); err != "" {
				return
			}
		}
	}
}

func (s *Service) GetRoleService(token string) (string, error) {
	return s.Repo.GetRoleRepo(token)
}

func (s *Service) GetCustomerByIDService(id int) (string, error) {
	return s.Repo.GetCustomerByIDRepo(id)
}
