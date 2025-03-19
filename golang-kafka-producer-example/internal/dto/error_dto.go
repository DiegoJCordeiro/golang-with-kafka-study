package dto

type ErrorDTO struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}
