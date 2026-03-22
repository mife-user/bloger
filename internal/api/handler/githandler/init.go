package githandler

import (
	"bloger/internal/domain"
)

// Handler 处理器
type GitHandler struct {
	Service domain.GitService
}

func NewGitHandler(service domain.GitService) *GitHandler {
	return &GitHandler{Service: service}
}
