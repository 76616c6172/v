package chat

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/sashabaranov/go-openai"
)

var Cmd = &Z.Cmd{
	Name:     `chat`,
	Summary:  `ask gpt-3.5`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		key, err := getAPIKey()
		if err != nil {
			fmt.Println("Error could not get OpenAI API key from file: ~/.config/v/chat")
			return err
		}

		response, err := CallLLM(strings.Join(args, " "), key)
		if err != nil {
			fmt.Println("Error getting response from OpenAI", err)
			return err
		}

		return renderMarkdown(response)
	},
}

func getAPIKey() (string, error) {
	homedir, err := os.UserHomeDir()
	file, err := os.Open(homedir + "/.config/v/chat")
	if err != nil {
		return "", err
	}
	defer file.Close()

	buf := new(strings.Builder)
	_, err = io.Copy(buf, file)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}

func CallLLM(p, k string) (string, error) {

	client := openai.NewClient(k)
	context := context.Background()

	request := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 1000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a large language model trained by OpenAI. Please follow the user's instructions carefully. Respond using markdown and be extremely concise.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: p,
			},
		},
		Stream: false,
	}

	response, err := client.CreateChatCompletion(context, request)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return response.Choices[0].Message.Content, err
}

func renderMarkdown(str string) error {
	output, err := glamour.Render(str, "dark")
	if err != nil {
		fmt.Println("Error rendering markdown text", err)
		return err
	}
	fmt.Print(output)

	return nil
}
