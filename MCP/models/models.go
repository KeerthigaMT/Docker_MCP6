package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// MonthModel represents the MonthModel schema from the OpenAPI specification
type MonthModel struct {
	Month int `json:"month,omitempty"`
}

// NamespaceMetadata represents the NamespaceMetadata schema from the OpenAPI specification
type NamespaceMetadata struct {
	Extrarepos []string `json:"extraRepos,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Datasets []DatasetModel `json:"datasets,omitempty"`
}

// WeekData represents the WeekData schema from the OpenAPI specification
type WeekData struct {
	Weeks []WeekModel `json:"weeks,omitempty"`
}

// MonthData represents the MonthData schema from the OpenAPI specification
type MonthData struct {
	Months []MonthModel `json:"months,omitempty"`
}

// WeekModel represents the WeekModel schema from the OpenAPI specification
type WeekModel struct {
	Week int `json:"week,omitempty"`
}

// UsersLoginRequest represents the UsersLoginRequest schema from the OpenAPI specification
type UsersLoginRequest struct {
	Password string `json:"password"` // The password or personal access token (PAT) of the Docker Hub account to authenticate with.
	Username string `json:"username"` // The username of the Docker Hub account to authenticate with.
}

// ResponseDataFile represents the ResponseDataFile schema from the OpenAPI specification
type ResponseDataFile struct {
	Size int64 `json:"size,omitempty"`
	Url string `json:"url,omitempty"`
}

// YearData represents the YearData schema from the OpenAPI specification
type YearData struct {
	Years []YearModel `json:"years,omitempty"`
}

// ResponseData represents the ResponseData schema from the OpenAPI specification
type ResponseData struct {
	Data []ResponseDataFile `json:"data,omitempty"`
}

// NamespaceData represents the NamespaceData schema from the OpenAPI specification
type NamespaceData struct {
	Namespaces []string `json:"namespaces,omitempty"`
}

// PostUsers2FALoginErrorResponse represents the PostUsers2FALoginErrorResponse schema from the OpenAPI specification
type PostUsers2FALoginErrorResponse struct {
	Detail string `json:"detail,omitempty"` // Description of the error.
}

// Users2FALoginRequest represents the Users2FALoginRequest schema from the OpenAPI specification
type Users2FALoginRequest struct {
	Code string `json:"code"` // The Time-based One-Time Password of the Docker Hub account to authenticate with.
	Login_2fa_token string `json:"login_2fa_token"` // The intermediate 2FA token returned from `/v2/users/login` API.
}

// PostUsersLoginSuccessResponse represents the PostUsersLoginSuccessResponse schema from the OpenAPI specification
type PostUsersLoginSuccessResponse struct {
	Token string `json:"token,omitempty"` // Created authentication token. This token can be used in the HTTP Authorization header as a JWT to authenticate with the Docker Hub APIs.
}

// PostUsersLoginErrorResponse represents the PostUsersLoginErrorResponse schema from the OpenAPI specification
type PostUsersLoginErrorResponse struct {
	Detail string `json:"detail"` // Description of the error.
	Login_2fa_token string `json:"login_2fa_token,omitempty"` // Short time lived token to be used on `/v2/users/2fa-login` to complete the authentication. This field is present only if 2FA is enabled.
}

// YearModel represents the YearModel schema from the OpenAPI specification
type YearModel struct {
	Year int `json:"year,omitempty"`
}

// DatasetModel represents the DatasetModel schema from the OpenAPI specification
type DatasetModel struct {
	Timespans []string `json:"timespans,omitempty"`
	Views []string `json:"views,omitempty"`
	Name string `json:"name,omitempty"`
}
