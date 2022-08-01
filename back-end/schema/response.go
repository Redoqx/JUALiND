package schema

type JSONResponse struct {
	Message string      `json:"msg"`
	Body    interface{} `json:"body,omitempty"`
	Error   string      `json:"error,omitempty"`
}
