package rpc_users

import (
	"database/sql"
	"errors"
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/internal/auth"
	http_errors "sms_portal/internal/errors"
	"sms_portal/internal/pagination"
	"sms_portal/internal/utils"
	"strconv"
)

func UsersIndex(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	// Authorize user
	user_id := r.Context().Value("user_id").(int64)

	if !auth.HasPermission(user_id, "users.list") {
		return nil, http_errors.Forbidden()
	}

	page, perPage, search := pagination.GetPaginationOptions(r)

	users, err := deps.Queries.ListUsers(r.Context(), sqlc.ListUsersParams{
		Search: search,
		Offset: int32((page - 1) * perPage),
		Limit:  int32(perPage),
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func UsersGet(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	// Extract path variable
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, http_errors.NotFound()
	}

	// Get User by ID
	user, err := deps.Queries.GetUserById(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, http_errors.NotFound()
		}
		return nil, http_errors.InternalServerError()
	}
	return user, nil
}
