package gitservice

import (
	"mifer/pkg/task"
	"mifer/pkg/utils"
	"context"
)

func (s *GitService) Save(ctx context.Context, token string) error {

	task.Do(ctx, func() error { // 加密token
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
	})

	return nil
}
