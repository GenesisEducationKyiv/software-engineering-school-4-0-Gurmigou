package rates

import "gorm.io/gorm"

type RateRepository interface {
	Where(query interface{}, args ...interface{}) RateRepository
	First(dest interface{}, conds ...interface{}) error
	Create(value interface{}) error
	Delete(value interface{}, conds ...interface{}) error
}

type GormRateRepository struct {
	db *gorm.DB
}

func NewRateRepository(db *gorm.DB) *GormRateRepository {
	return &GormRateRepository{db: db}
}

func (r *GormRateRepository) Where(query interface{}, args ...interface{}) RateRepository {
	r.db = r.db.Where(query, args...)
	return r
}

func (r *GormRateRepository) First(dest interface{}, conds ...interface{}) error {
	return r.db.First(dest, conds...).Error
}

func (r *GormRateRepository) Create(value interface{}) error {
	return r.db.Create(value).Error
}

func (r *GormRateRepository) Delete(value interface{}, conds ...interface{}) error {
	return r.db.Delete(value, conds...).Error
}
