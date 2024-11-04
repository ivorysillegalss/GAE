package repository

import (
	"database/sql"
	"gae-backend-web/domain"
	"gae-backend-web/infrastructure/hive"
	"gae-backend-web/infrastructure/log"
)

type userRepository struct {
	hive hive.Client
}

func NewUserRepository(hc hive.Client) domain.UserRepository {
	return &userRepository{hive: hc}
}

func (u *userRepository) QueryUser(username string) *[]*domain.Contributor {
	//TODO query条件补全
	query, err := u.hive.Query("")
	if err != nil {
		log.GetTextLogger().Warn("error querying user, username is: " + username)
		return nil
	}
	var c []*domain.Contributor
	for query.Next() {
		cv, err := newContributorValue(query)
		if err != nil {
			if err != nil {
				log.GetTextLogger().Warn("error querying user, username is: " + username)
				return nil
			}
		}
		c = append(c, cv)
	}
	return &c
}

// 根据查询结果构建并返回 Contributor 实例
func newContributorValue(rows *sql.Rows) (*domain.Contributor, error) {
	var contributor domain.Contributor
	err := rows.Scan(
		&contributor.Login,
		&contributor.Id,
		&contributor.AvatarURL,
		&contributor.URL,
		&contributor.HTMLURL,
		&contributor.GravatarID,
		&contributor.Name,
		&contributor.Company,
		&contributor.Blog,
		&contributor.Location,
		&contributor.Email,
		&contributor.Bio,
		&contributor.PublicRepos,
		&contributor.PublicGists,
		&contributor.FollowerCount,
		&contributor.FollowingCount,
		&contributor.CreatedAt,
		&contributor.UpdatedAt,
		&contributor.Type,
	)
	if err != nil {
		return nil, err
	}
	return &contributor, nil
}
