package link

import (
	"go-advance/pkg/db"

	"gorm.io/gorm"
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

func (r *LinkRepository) Update(link *Link) error {
	result := r.DB.Clauses(clause.Returning{}).UpdateColumns(link)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LinkRepository) Delete(id uint) error {
	link := Link{Model: gorm.Model{ID: id}}
	result := r.DB.Delete(&link)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := r.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil
}

func (r *LinkRepository) Count() int64 {
	var count int64

	r.DB.Table("links").Where("deleted_at is null").Count(&count)

	return count
}

func (r *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link

	r.DB.
		Table("links").
		Where("deleted_at is null").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}
