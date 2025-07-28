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

func (r *LinkRepository) Create(link *Link) {

}
