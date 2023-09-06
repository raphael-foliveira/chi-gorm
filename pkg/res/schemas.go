package res

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
