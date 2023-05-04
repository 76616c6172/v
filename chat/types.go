package chat

const CHAT_GPT_URL = "https://api.openai.com/v1/chat/completions"
const OPENAI_MODEL = "gpt-3.5-turbo"

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type PostBody struct {
	Model    string    `json:"model"`
	Messages []message `json:"messages"`
}

type ResponseBody struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
