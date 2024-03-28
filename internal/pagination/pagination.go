package pagination

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type Paginator struct {
	CurrentPage int              `json:"current_page"`
	From        int              `json:"from"`
	LastPage    int              `json:"last_page"`
	Links       []PaginationLink `json:"links"`
	Path        string           `json:"path"`
	PerPage     int              `json:"per_page"`
	To          int              `json:"to"`
	Total       int              `json:"total"`
}

type PaginatedResults struct {
	Data  interface{}     `json:"data"`
	Links PaginationLinks `json:"links"`
	Meta  Paginator       `json:"meta"`
}

func NewPaginatedResults(options ...PaginatedResultsOption) *PaginatedResults {
	p := &PaginatedResults{}

	for _, option := range options {
		option(p)
	}

	return p
}

type PaginationLinks struct {
	First *string `json:"first"`
	Last  *string `json:"last"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
}

type PaginationLink struct {
	Url    *string `json:"url"`
	Label  *string `json:"label"`
	Active bool    `json:"active"`
}

type PaginatedResultsOption func(*PaginatedResults)

func WithData(data interface{}) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Data = data
	}
}

func WithPerPage(perPage int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.PerPage = perPage
	}
}

func WithPage(page int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.CurrentPage = page
	}
}

func WithFrom(from int) PaginatedResultsOption {
	fmt.Printf("Setting From to: %d\n", from)
	return func(p *PaginatedResults) {
		p.Meta.From = from
	}
}

func WithTo(to int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.To = to
	}
}

func WithTotal(total int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.Total = total
	}
}

func WithLastPage(lastPage int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.LastPage = lastPage
	}
}

func GetPaginationFieldsFromRequest(r *http.Request) (int, int, sql.NullString) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
	if err != nil {
		perPage = 10
	}

	search := r.URL.Query().Get("search")
	var sqlSearch sql.NullString
	if search == "" {
		sqlSearch = sql.NullString{}
	} else {
		sqlSearch = sql.NullString{
			String: "%" + search + "%",
			Valid:  true,
		}
	}

	return page, perPage, sqlSearch
}
