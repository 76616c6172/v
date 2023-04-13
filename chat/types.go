package chat

type Prompt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type APIRequest struct {
	Messages []Prompt `json:"messages"`
}

type APIResponse struct {
	Choices []struct {
		Message Prompt `json:"message"`
	} `json:"choices"`
}
