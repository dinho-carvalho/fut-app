package database

import (
	"context"
)

type QueryOptions struct {
	Page     int
	PageSize int
	OrderBy  string
	Search   string
}

type PaginatedResult struct {
	Data       interface{}
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}

// Paginate é um método auxiliar para paginação
func (db *Database) Paginate(ctx context.Context, model interface{}, opts QueryOptions) (*PaginatedResult, error) {
	var total int64
	query := db.WithContext(ctx).Model(model)

	if opts.Search != "" {
		// TODO implement search logic
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if opts.OrderBy != "" {
		query = query.Order(opts.OrderBy)
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}
	if opts.PageSize <= 0 {
		opts.PageSize = 10
	}

	var data interface{}
	err := query.Offset((opts.Page - 1) * opts.PageSize).
		Limit(opts.PageSize).
		Find(&data).Error
	if err != nil {
		return nil, err
	}

	totalPages := int(total) / opts.PageSize
	if int(total)%opts.PageSize != 0 {
		totalPages++
	}

	return &PaginatedResult{
		Data:       data,
		Total:      total,
		Page:       opts.Page,
		PageSize:   opts.PageSize,
		TotalPages: totalPages,
	}, nil
}
