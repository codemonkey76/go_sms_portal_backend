package routes

import (
	"net/http"
	"sms_portal/routes/api"
	"sms_portal/utils"
)

type RouteRegistrar struct {
	Mux *http.ServeMux
}

func NewRouteRegistrar(mux *http.ServeMux) *RouteRegistrar {
	return &RouteRegistrar{Mux: mux}
}

func (rr *RouteRegistrar) AddHandler(method, prefix, route string, handlerFunc utils.HandlerFunc) {
	fullPath := method + " " + prefix + route
	handler := utils.HandleRequest(handlerFunc)

	rr.Mux.HandleFunc(fullPath, handler)
}

func RegisterRoutes(mux *http.ServeMux) {
	rr = NewRouteRegistrar(mux)
	api.RegisterRoutes("api/v1", rr)
}
