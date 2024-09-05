package news_banner

import "backend/internal/database"

type NewsBannerEntity struct {
	db *database.StorageDb
}

func NewNewsBannerEntity(db *database.StorageDb) *NewsBannerEntity {
	return &NewsBannerEntity{
		db: db,
	}
}
