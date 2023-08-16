package utils

import "gorm.io/gorm"

type paginationConfig struct {
	perPage     int
	totalPages  int
	totalItems  int
	currentPage int
}

func NewPaginationConfig(perPage, totalItems, currentPage int) *paginationConfig {
	return &paginationConfig{
		perPage:     perPage,
		totalPages:  totalItems / perPage,
		totalItems:  totalItems,
		currentPage: currentPage,
	}
}

func (pg *paginationConfig) Paginate(tx *gorm.DB) *gorm.DB {
	return tx.Offset((pg.currentPage - 1) * pg.perPage).Limit(pg.perPage)
}
