package rpc_users

import (
	"database/sql"
	"errors"
	"math"
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

	page, perPage, search := pagination.GetPaginationFieldsFromRequest(r)

	users, err := deps.Queries.ListUsers(r.Context(), sqlc.ListUsersParams{
		Search: search,
		Offset: int32((page - 1) * perPage),
		Limit:  int32(perPage),
	})
	if err != nil {
		return nil, err
	}

	totalCount := 0
	from := 0
	to := 0
	if len(users) > 0 {
		totalCount = int(users[0].TotalCount)
		from = ((page - 1) * perPage) + 1
		to = len(users) + ((page - 1) * perPage)
	}

	paginator := pagination.NewPaginatedResults(
		pagination.WithData(users),
		pagination.WithPerPage(perPage),
		pagination.WithPage(page),
		pagination.WithFrom(from),
		pagination.WithTo(to),
		pagination.WithTotal(totalCount),
		pagination.WithLastPage(int(math.Ceil(float64(totalCount)/float64(perPage)))),
	)

	return paginator, nil
}

func UsersGet(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	// Extract path variable
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, http_errors.NotFound()
	}

	// Get User by ID
	user, err := deps.Queries.GetUserById(r.Context(), int64(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, http_errors.NotFound()
		}
		return nil, http_errors.InternalServerError()
	}
	return user, nil
}
