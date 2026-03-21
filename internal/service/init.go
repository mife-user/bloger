package service

import "bloger/internal/domain"

type Service struct {
	Repo domain.Repo
}

func NewService(repo domain.Repo) *Service {
	return &Service{Repo: repo}
}
