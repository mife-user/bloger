package gitrepo

import (
	"bloger/pkg/conf"
	"bloger/pkg/db"
)

type GitRepo struct {
	db *db.JSONFileDB
}

func NewGitRepo(c *conf.Config) *GitRepo {
	return &GitRepo{
		db: db.NewJSONFileDB(c.Git.LockPath),
	}
}
