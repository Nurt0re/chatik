package service

import (
	"github.com/Nurt0re/chatik"
	"github.com/Nurt0re/chatik/pkg/repository"
)



type UpdaterService struct {
    repo repository.Updater

}

func NewUpdater(repo repository.Updater) *UpdaterService {
    return &UpdaterService{repo: repo}
}

func (s *UpdaterService) UpdateUser(id int, input chatik.User) error {

    return s.repo.UpdateUser(id, input)
}

func (s *UpdaterService) DeleteUser(id int) error {
	
	return s.repo.DeleteUser(id)
}
func (s *UpdaterService) GetUser(id int) (chatik.User, error) {
	
	return s.repo.GetUser(id)
}
func (s *UpdaterService) GetAllUsers() ([]chatik.User, error) {
	
	return s.repo.GetAllUsers()
}