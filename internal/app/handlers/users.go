package handlers

import (
	"net/http"
	"sms_portal/internal/utils"
)

type UserLoginHandler struct {
}

func (u *UserLoginHandler) Validate(r *http.Request) error {
	return nil
}

func (u *UserLoginHandler) Authorize(r *http.Request) bool {
	return true
}

func (u *UserLoginHandler) Execute(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	return nil, nil
}
