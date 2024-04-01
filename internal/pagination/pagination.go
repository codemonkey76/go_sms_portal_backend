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

func NewPaginationLink(url, label *string, active bool) PaginationLink {
	return PaginationLink{
		Url:    url,
		Label:  label,
		Active: active,
	}
}

type PaginatedResultsOption func(*PaginatedResults)

func WithPath(path string) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		p.Meta.Path = path
	}
}

func SetupLinks(linksEachSide int) PaginatedResultsOption {
	return func(p *PaginatedResults) {
		first := p.Meta.Path + "?page=1"
		last := p.Meta.Path + "?page=" + strconv.Itoa(p.Meta.LastPage)
		p.Links.First = &first
		p.Links.Last = &last

		if p.Meta.LastPage > p.Meta.CurrentPage {
			// Setup next link
			next := p.Meta.Path + "?page=" + strconv.Itoa(p.Meta.CurrentPage+1)
			p.Links.Next = &next
		}

		if p.Meta.CurrentPage > 1 {
			// Setup prev link
			prev := p.Meta.Path + "?page=" + strconv.Itoa(p.Meta.CurrentPage-1)
			p.Links.Prev = &prev
		}

		numLeader := 2
		p.Meta.Links = getPageLinks(p.Meta.Path, p.Meta.CurrentPage, p.Meta.LastPage, linksEachSide, numLeader)
	}
}

func getPageLinks(path string, currentPage, totalPages, eachSide, numLeader int) []PaginationLink {
	links := make([]PaginationLink, 0)

	// Previous
	{
		label := "&laquo; Previous"
		var url *string = nil
		active := false

		l := NewPaginationLink(url, &label, active)
		links = append(links, l)
	}

	// Add all pages
	hasNotManyPages := totalPages <= (3 + (eachSide * 2) + (numLeader * 2))
	isNearStart := currentPage <= (3 + eachSide + numLeader)
	//_isNearEnd := currentPage >= (totalPages - (2 + (eachSide * 2)))

	if hasNotManyPages {
		fmt.Printf("Not many pages\n")
		for i := 1; i <= totalPages; i++ {
			url := fmt.Sprintf("%s?page=%d", path, i)
			active := i == currentPage
			label := fmt.Sprintf("%d", i)
			l := NewPaginationLink(&url, &label, active)
			links = append(links, l)
		}
	} else {
		fmt.Printf("Has a lot of pages\n")
		if isNearStart {
			fmt.Printf("Is near start\n")
			for i := 1; i <= currentPage+eachSide; i++ {
				fmt.Printf("Adding link for page %d\n", i)
				url := fmt.Sprintf("%s?page=%d", path, i)
				active := i == currentPage
				label := fmt.Sprintf("%d", i)
				l := NewPaginationLink(&url, &label, active)
				links = append(links, l)
			}
			// Add ...
			{
				var url *string = nil
				active := false
				label := "..."
				l := NewPaginationLink(url, &label, active)
				links = append(links, l)
			}

			// Add the last 2 pages
			for i := totalPages - 1; i <= totalPages; i++ {
				url := fmt.Sprintf("%s?page=%d", path, i)
				active := i == currentPage
				label := fmt.Sprintf("%d", i)
				l := NewPaginationLink(&url, &label, active)
				links = append(links, l)
			}
		} else {
			fmt.Printf("Is not near start\n")
			for i := 1; i <= 2; i++ {
				url := fmt.Sprintf("%s?page=%d", path, i)
				active := false
				label := fmt.Sprintf("%d", i)
				l := NewPaginationLink(&url, &label, active)
				links = append(links, l)
			}

			for i := currentPage - eachSide; i <= currentPage+eachSide; i++ {
				url := fmt.Sprintf("%s?page=%d", path, i)
				active := i == currentPage
				label := fmt.Sprintf("%d", i)
				l := NewPaginationLink(&url, &label, active)
				links = append(links, l)
			}
		}

	}

	// Next
	{
		label := "Next &raquo;"
		var url *string = nil
		active := false

		l := NewPaginationLink(url, &label, active)
		links = append(links, l)
	}
	return links
}

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
