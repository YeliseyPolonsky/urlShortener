package link

import (
	"go-advance/pkg/db"

	"gorm.io/gorm/clause"
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

func (r *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := r.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *LinkRepository) IsExist(name string, value string) bool {
	var link Link
	r.DB.First(&link, name+" = ?", value)

	return link.ID != 0
}

func (r *LinkRepository) Update(link *Link) error {
	result := r.DB.Clauses(clause.Returning{}).UpdateColumns(link)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
