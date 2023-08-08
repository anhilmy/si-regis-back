package response

type FailedBadRequest struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
