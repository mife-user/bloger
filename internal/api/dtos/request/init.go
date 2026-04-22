package request

type ChatRequest struct {
	Message map[string]any `json:"message"`
}
