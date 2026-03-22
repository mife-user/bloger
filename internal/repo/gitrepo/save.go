package gitrepo

import (
	"bloger/internal/model/gitmodel"
)

func (r *GitRepo) Save(token string) error {
	model := gitmodel.GitModel{Token: token}
	return r.db.Save(model)
}
