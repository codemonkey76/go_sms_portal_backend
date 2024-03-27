package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"sms_portal/db/sqlc"
	http_errors "sms_portal/internal/errors"
	"sms_portal/internal/middleware"
	"sms_portal/internal/ui"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, HandlerDependencies) (interface{}, error)

func HandleRequest(deps HandlerDependencies, handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(w, r, deps)
		if err != nil {
			var httpErr *http_errors.HttpError
			if errors.As(err, &httpErr) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(httpErr.Code)
				errorMessage := map[string]string{"error": httpErr.Message}
				json.NewEncoder(w).Encode(errorMessage)
				ui.Error(httpErr.Message)
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
	Middleware  middleware.Middleware
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

func (rr *RouteRegistrar) AddHandler(method, prefix, route string, handlerFunc HandlerFunc, middleware middleware.Middleware) {
	fullPath := method + " " + prefix + route
	handler := HandleRequest(rr.Deps, handlerFunc)

	if middleware != nil {
		middlewareHandler := middleware(http.HandlerFunc(handler))
		rr.Mux.Handle(fullPath, middlewareHandler)
	} else {
		rr.Mux.Handle(fullPath, http.HandlerFunc(handler))
	}
	newRoute := Route{
		Method:      method,
		Prefix:      prefix,
		Route:       route,
		HandlerFunc: handlerFunc,
		Middleware:  middleware,
	}
	rr.Routes = append(rr.Routes, newRoute)
}
