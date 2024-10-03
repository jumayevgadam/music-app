package pagination

import (
	"fmt"
	"math"
	"strconv"

	"github.com/jumayevgadam/music-app/pkg/errlst"
	"github.com/labstack/echo/v4"
)

const (
	defaultSize = 10
)

// Compile time check to ensure PaginationQuery implements Pagination interface
var _ PaginationOps = (*PaginationQuery)(nil)

// Pagination interface keeps needed methods for pagination
type PaginationOps interface {
	SetSize(sizeQuery string) error
	SetPage(pageQuery string) error
	SetOrderBy(orderByQuery string)
	GetOffset() int
	GetLimit() int
	GetOrderBy() string
	GetPage() int
	GetSize() int
	GetQueryString() string
}

// PaginationQuery params
type PaginationQuery struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

// Set Page size, SetSize
func (q *PaginationQuery) SetSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.Size = defaultSize
		return nil
	}

	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return errlst.ParseErrors(err)
	}
	q.Size = n

	return nil
}

// Set Page number, SetPage
func (q *PaginationQuery) SetPage(pageQuery string) error {
	if pageQuery == "" {
		q.Size = 0
		return nil
	}

	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return errlst.ParseErrors(err)
	}
	q.Page = n

	return nil
}

// Set order by, SetOrderBy
func (q *PaginationQuery) SetOrderBy(orderByQuery string) {
	q.OrderBy = orderByQuery
}

// GetOffset is
func (q *PaginationQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}

	return (q.Page - 1) * q.Size
}

// GetLimit is
func (q *PaginationQuery) GetLimit() int {
	return q.Size
}

// GetOrderBy is
func (q *PaginationQuery) GetOrderBy() string {
	return q.OrderBy
}

// GetPage is
func (q *PaginationQuery) GetPage() int {
	return q.Page
}

// GetSize is
func (q *PaginationQuery) GetSize() int {
	return q.Size
}

// GetQueryString is
func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", q.Page, q.Size, q.OrderBy)
}

// GetPaginationFromCtx is
func GetPaginationFromCtx(c echo.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}

	// set page
	if err := q.SetPage(c.QueryParam("page")); err != nil {
		return nil, errlst.ParseErrors(err)
	}

	// set size
	if err := q.SetSize(c.QueryParam("size")); err != nil {
		return nil, errlst.ParseErrors(err)
	}
	// set orderby
	q.SetOrderBy(c.QueryParam("orderBy"))

	return q, nil
}

// GetTotalPages is
func GetTotalPages(totalCount int, pageSize int) int {
	d := float64(totalCount) / float64(pageSize)
	return int(math.Ceil(d))
}

// GetHasMore is
func GetHasMore(currentPage int, totalCount int, pageSize int) bool {
	return currentPage < totalCount/pageSize
}
