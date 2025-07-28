package link

import (
	"go-advance/pkg/db"
)

type LinkRepository struct {
	*db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Db: database,
	}
}

func (r *LinkRepository) Create(link *Link) error {
	result := r.DB.Create(link)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
