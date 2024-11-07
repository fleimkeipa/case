package controller

type FailureResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
