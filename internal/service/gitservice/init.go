package gitservice

import (
	"bloger/internal/repo/gitrepo"
)

type GitService struct {
	Repo gitrepo.GitRepo
}

func NewGitService(repo *gitrepo.GitRepo) *GitService {
	return &GitService{Repo: *repo}
}
