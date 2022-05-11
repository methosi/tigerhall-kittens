package storage

import "time"

// PaginationFilters allows specifying page and page size when
// querying data.
type PaginationFilters struct {
	// Page allows specifying the page of results to return.
	// If Page <= 0, the first page will be returned.
	Page int

	// PageSize allows specifying the specifying number of results to return.
	// If PageSize <= 0, the default page size of the implementation will be used.
	PageSize int
}

// DateFilters allows specifying start and end time when
// querying data.
type DateFilters struct {
	StartDate, EndDate time.Time
}

type Filters struct {
	// filter by ID
	IDs []string

	// filter by start and end date
	DateFilters

	// pagination (filter by page and page size)
	PaginationFilters
}

type FilterOption func(*Filters)

func IDs(ids []string) FilterOption {
	return func(f *Filters) {
		f.IDs = ids
	}
}

func StartDate(date time.Time) FilterOption {
	return func(f *Filters) {
		f.StartDate = date
	}
}

func EndDate(date time.Time) FilterOption {
	return func(f *Filters) {
		f.EndDate = date
	}
}

func Page(page int) FilterOption {
	return func(f *Filters) {
		f.Page = page
	}
}

func PageSize(pageSize int) FilterOption {
	return func(f *Filters) {
		f.PageSize = pageSize
	}
}

// OrderBy represents an order for returning query results.
type OrderBy int8

const (
	OrderByDescending OrderBy = iota
	OrderByAscending
)

// Ordering database
func (o OrderBy) String() string {
	switch o {
	case OrderByDescending:
		return "DESC"
	case OrderByAscending:
		return "ASC"
	default:
		return ""
	}
}
