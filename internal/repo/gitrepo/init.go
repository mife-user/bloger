package gitrepo

import (
	"mifer/pkg/conf"
	"mifer/pkg/db"
)

type GitRepo struct {
	db *db.JSONFileDB
}

func NewGitRepo(c *conf.Config) *GitRepo {
	return &GitRepo{
		db: db.NewJSONFileDB(c.Git.LockPath),
	}
}
