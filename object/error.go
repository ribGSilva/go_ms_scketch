package object

type errorResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func NewErrorResponse(code int, message string, details ...string) *errorResponse {
	return &errorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}
}
