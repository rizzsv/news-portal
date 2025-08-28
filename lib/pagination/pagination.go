package pagination

import (
	"math"
	"news-portal/internal/core/domain/entity"
)

type PaginationInterface interface {
	AddPagination(totalData, page, perpage int) (*entity.Page, error)
}

type Options struct{}

// AddPagination implements PaginationInterface.
func (o *Options) AddPagination(totalData int, page int, perpage int) (*entity.Page, error) {
	newPage := page

	if newPage <= 0 {
		return nil, ErrorPage
	}

	limitData := 10
	if perpage > 0 {
		limitData = perpage
	}

	totalPage  := int(math.Ceil(float64(totalData) / float64(limitData)))

	last := (newPage * limitData)
	first := last - totalData

	if totalData < last {
		last = totalData
	}

	zeroPage := &entity.Page{PageCount: 1, Page: newPage}
	if totalData == 0 && newPage == 1 {
		return zeroPage, nil
	}

	if newPage > totalPage {
		return nil, ErrorMaxPage
	}

	pages := &entity.Page{
		Page:       newPage,
		PerPage:    perpage,
		PageCount:  totalPage,
		TotalCount: totalData,
		First:      first,
		Last:       last,
	}
	return pages, nil

}

func NewPagination() PaginationInterface {
	pagination := new(Options)

	return pagination
}
