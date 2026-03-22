package gitservice

import "bloger/pkg/utils"

func (s *GitService) Save(token string) error {
	// 加密token
	hash, err := utils.HashPassword(token)
	if err != nil {
		return err
	}
	// 保存token到数据库
	err = s.Repo.Save(hash)
	if err != nil {
		return err
	}
	return nil
}
