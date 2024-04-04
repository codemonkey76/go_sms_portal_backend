package handlers

import (
	"net/http"
	"sms_portal/internal/utils"
)

type Handler interface {
	Validate(r *http.Request) error
	Authorize(r *http.Request) bool
	Execute(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error)
}

