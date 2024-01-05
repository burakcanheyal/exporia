package repository

import (
	"exporia/internal/domain/dto"
	"exporia/internal/domain/enum"
	"gorm.io/gorm"
)

func Paginate(pagination dto.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Offset((pagination.Page - 1) * enum.PaginationLimit).Limit(enum.PaginationLimit)
	}
}
