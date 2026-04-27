package domain

type ChatRequest struct {
	Content string `json:"content"`
}

type ChatResponse struct {
	Content string `json:"content"`
}
