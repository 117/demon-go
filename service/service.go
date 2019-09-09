package service

import (
	"github.com/takama/daemon"
)

// Service -
type Service struct {
	ID string `json:"id"`
}

// Spawn -
func (s *Service) Spawn() error {
	service, err := daemon.New("id goes here", "default demon description")
	if err != nil {
		return err
	}
	if _, err = service.Remove(); err != nil {
		return err
	}
	return err
}
