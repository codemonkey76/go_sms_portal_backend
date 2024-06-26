package rpc_auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/internal/config"
	"sms_portal/internal/errors"
	"sms_portal/internal/ui"
	"sms_portal/internal/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func CheckSession(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	w.WriteHeader(http.StatusOK)
	response := CheckSessionResponse{LoggedIn: true}

	return response, nil
}

type UserLoginForm struct {
	Email    string `validate:"required,email"`
	Password string `valdiate:"required"`
}

func validate(formData *interface{}, body io.ReadCloser) {
	validate := validator.New()
}

func AuthLogin(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	validate := validator.New()
	creds := UserLoginForm{
		Email:    "",
		Password: "",
	}

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		ui.Info("Error decoding request body: " + err.Error())
		return nil, errors.InvalidCredentials()
	}

	err = validate.Struct(creds)

	ui.Info(fmt.Sprintf("Received Creds: %s", creds))
	user, err := deps.Queries.GetUserByEmail(r.Context(), creds.Email)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	}

	if err != nil {
		ui.Info("Error logging in: " + err.Error())
		return nil, errors.InvalidCredentials()
	}
	ui.Info("User logged in: " + creds.Email)

	tokenString, err := utils.GenerateSessionToken()

	if err != nil {
		ui.Info("Error generating session token: " + err.Error())
		return nil, errors.InternalServerError()
	}

	_, err = deps.Queries.CreateSession(r.Context(), sqlc.CreateSessionParams{
		ID:           tokenString,
		IpAddress:    sql.NullString{String: r.RemoteAddr, Valid: true},
		UserAgent:    sql.NullString{String: r.UserAgent(), Valid: true},
		Payload:      "",
		LastActivity: time.Now().Unix(),
		UserID:       int64(user.ID),
	})

	if err != nil {
		ui.Info("Error creating session: " + err.Error())
		return nil, errors.InternalServerError()
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    tokenString,
		Path:     "/",
		Expires:  time.Now().Add(config.SessionExpiration * time.Minute),
		Secure:   true,
		HttpOnly: true,
	})

	return_user := sqlc.User{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Password:        "",
		Active:          false,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}

	permissions, err := deps.Queries.ListUserPermissions(r.Context(), user.ID)

	if err != nil {
		permissions = []string{}
	}
	response := LoginResponse{
		Success: true,
		Message: "Login successful.",
		Data: UserData{
			User:        return_user,
			Permissions: permissions,
		},
	}

	return response, nil
}

type CheckSessionResponse struct {
	LoggedIn bool `json:"loggedIn"`
}

type LoginResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Data    UserData `json:"data"`
}

type UserData struct {
	User        sqlc.User `json:"user"`
	Permissions []string  `json:"permissions"`
}

func AuthLogout(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	user_id := r.Context().Value("user_id").(int64)
	err := deps.Queries.DeleteSessionByUserId(r.Context(), user_id)
	if err != nil {
		return nil, errors.InternalServerError()
	}

	return nil, nil
}

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
