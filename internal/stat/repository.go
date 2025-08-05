package stat

import (
	"go-advance/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{
		Db: database,
	}
}

func (r *StatRepository) AddClick(linkId uint) {
	var stat Stat
	date := datatypes.Date(time.Now())
	r.DB.Find(&stat, "link_id = ? and date = ?", linkId, date)
	if stat.ID == 0 {
		stat = Stat{
			LinkID: linkId,
			Clicks: 1,
			Date:   date,
		}
		r.DB.Create(&stat)
	} else {
		stat.Clicks++
		r.DB.Save(&stat)
	}
}

func (r *StatRepository) GetStat(by string, from, to time.Time) []GetStatRespone {
	var stats []GetStatRespone
	var selectQuery string
	switch by {
	case GroupByDay:
		selectQuery = "to_char(date,'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date,'YYYY-MM') as period, sum(clicks)"
	}
	r.DB.Table("stats").
		Select(selectQuery).
		Where("date between ? and ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)

	return stats
}
