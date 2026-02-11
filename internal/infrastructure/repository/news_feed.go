package repository

import (
	mNewsFeed "clone-instagram-service/internal/domain/model/news_feed"
	aNewsFeed "clone-instagram-service/internal/domain/model/news_feed/aggregate"
	eNewsFeed "clone-instagram-service/internal/domain/model/news_feed/entity"
	"context"
	"log"

	"gorm.io/gorm"
)

type newsFeedRepository struct {
	gormDB *gorm.DB
}

func NewNewsFeedRepository(gormDB *gorm.DB) *newsFeedRepository {
	return &newsFeedRepository{gormDB: gormDB}
}

var _ mNewsFeed.NewsFeedRepository = (*newsFeedRepository)(nil)

func (r *newsFeedRepository) GetNewsFeedByFilter(ctx context.Context, filter eNewsFeed.NewsFeedFilter) ([]aNewsFeed.NewsFeed, error) {
	newsFeeds := []aNewsFeed.NewsFeed{}
	err := r.gormDB.Raw(`
	select 
		m.id as media_id,
		m.owner_user_id as media_owner_user_id,
		m.file_storage_link as media_file_storage_link,
		m.caption as media_caption,
		m.created_at as media_created_at,
		u.id as user_id,
		u.google_sub_id as user_sub,
		u.name as user_name,
		u.given_name as user_given_name,
		u.family_name as user_family_name,
		u.picture as user_picture,
		u.email as user_email,
		u.created_at as user_created_at
	from
		followings f
	right join medias m on
		m.owner_user_id = f.target_user_id
	inner join users u on
		u.id = f.target_user_id
	where
		f.user_id = ?
		offset ?
		limit ?
	`, filter.UserID, filter.Offset, filter.Limit).Scan(&newsFeeds).Error

	if err != nil {
		log.Printf("Error while inserting media metadata into database. Here's why: %v\n", err)
		return newsFeeds, err
	}

	return newsFeeds, nil
}
