package gitservice

import (
	"mifer/internal/repo/gitrepo"
)

type GitService struct {
	Repo *gitrepo.GitRepo
}

func NewGitService(repo *gitrepo.GitRepo) *GitService {
	return &GitService{Repo: repo}
}
