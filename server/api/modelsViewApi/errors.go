package modelsViewApi

type BadRequestError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
