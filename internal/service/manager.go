package service

import "errors"

type Service struct {
}

func New(repo *repository.Repository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("NO repo")
	}

	return &Service{}, nil
}
