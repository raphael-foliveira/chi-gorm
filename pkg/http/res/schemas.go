package res

type ApiError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}
