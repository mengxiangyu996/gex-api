package query

import "gorm.io/gorm"

// 条件
type Query struct{}

func WithId(id int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}
