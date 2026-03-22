package domain

type GitService interface {
	Save(token string) error
}

type GitRepo interface {
	Save(token string) error
}
