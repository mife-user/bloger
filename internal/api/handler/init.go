package handler

import (
	"bloger/internal/domain"
)

// Handler 处理器
type Handler struct {
	Service domain.Service
}

func NewHandler(service domain.Service) *Handler {
	return &Handler{Service: service}
}
