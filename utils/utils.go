package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"sms_portal/db/sqlc"
	httperrors "sms_portal/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, HandlerDependencies) (interface{}, error)

func HandleRequest(deps HandlerDependencies, handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %-8s - %s", r.Method, r.URL.RequestURI())
		data, err := handler(w, r, deps)
		if err != nil {
			var httpErr *httperrors.HttpError
			if errors.As(err, &httpErr) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(httpErr.Code)
				errorMessage := map[string]string{"error": httpErr.Message}
				json.NewEncoder(w).Encode(errorMessage)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		response, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

type RouteRegistrar struct {
	Mux    *http.ServeMux
	Routes []Route
	Deps   HandlerDependencies
}

type HandlerDependencies struct {
	DB      *sql.DB
	Queries *sqlc.Queries
}

type Route struct {
	Method      string
	Prefix      string
	Route       string
	HandlerFunc HandlerFunc
}

type RouteRegistrarOption func(*RouteRegistrar)

func WithDB(db *sql.DB, queries *sqlc.Queries) RouteRegistrarOption {
	return func(rr *RouteRegistrar) {
		rr.Deps = HandlerDependencies{
			DB:      db,
			Queries: queries,
		}
	}
}

func NewRouteRegistrar(mux *http.ServeMux, opts ...RouteRegistrarOption) *RouteRegistrar {
	rr := &RouteRegistrar{Mux: mux}

	for _, opt := range opts {
		opt(rr)
	}

	return rr
}

func (rr *RouteRegistrar) AddHandler(method, prefix, route string, handlerFunc HandlerFunc) {
	fullPath := method + " " + prefix + route
	handler := HandleRequest(rr.Deps, handlerFunc)

	newRoute := Route{
		Method:      method,
		Prefix:      prefix,
		Route:       route,
		HandlerFunc: handlerFunc,
	}
	rr.Routes = append(rr.Routes, newRoute)

	rr.Mux.HandleFunc(fullPath, handler)
}

type StringPart struct {
	Part    string
	IsField bool
}

func GetPartsByField(input string) []StringPart {
	re := regexp.MustCompile(`{[^}]*}`)

	matches := re.FindAllStringIndex(input, -1)
	var parts []StringPart
	lastIndex := 0
	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			parts = append(parts, StringPart{Part: input[lastIndex:start], IsField: false})
		}

		parts = append(parts, StringPart{Part: input[start:end], IsField: true})

		lastIndex = end
	}

	if lastIndex < len(input) {
		parts = append(parts, StringPart{Part: input[lastIndex:], IsField: false})
	}

	return parts
}
