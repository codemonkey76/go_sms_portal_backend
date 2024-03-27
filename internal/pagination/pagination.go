package pagination

import (
	"database/sql"
	"net/http"
	"strconv"
)

type Paginator struct {
	CurrentPage  int              `json:"current_page"`
	Data         interface{}      `json:"data"`
	FirstPageUrl string           `json:"first_page_url"`
	From         int              `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageUrl  string           `json:"last_page_url"`
	Links        []PaginationLink `json:"links"`
	NextPageUrl  string           `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      int              `json:"per_page"`
	PrevPageUrl  string           `json:"prev_page_url"`
	To           int              `json:"to"`
	Total        int              `json:"total"`
}

type PaginationLink struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

func GetPaginationOptions(r *http.Request) (int, int, sql.NullString) {
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
