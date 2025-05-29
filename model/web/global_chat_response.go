package web

type GlobalChatResponse struct {
    AnonId    string `json:"anon_id"`
    Message   string `json:"message"`
    CreatedAt string `json:"created_at"`
}