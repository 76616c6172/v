package chat

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	Z "github.com/rwxrob/bonzai/z"

	"github.com/76616c6172/help"
	"github.com/charmbracelet/glamour"
	"github.com/theckman/yacspin"
)

var Cmd = &Z.Cmd{
	Name:    `chat`,
	Summary: `interactively chat with gpt-3.5`,
	Description: `
		Reads OpenAI API Key from $HOME/.config/v/chat/key and system prompt from $HOME/.config/v/chat/sysprompt
		`,

	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		sysprompt, err := getStringFromFile("sysprompt")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		apikey, err := getStringFromFile("key")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		run_chat_in_terminal(sysprompt, apikey)

		return nil
	},
}

func run_chat_in_terminal(sysprompt, apiKey string) {
	var messages []message
	messages = append(messages, message{
		Role:    "system",
		Content: sysprompt,
	})

	for {
		//print("▸ ")
		print("⌗ ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println("")
			break
		}

		name := scanner.Text()

		messages = append(messages, message{
			Role:    "user",
			Content: name,
		})

		cfg := yacspin.Config{
			Frequency:       100 * time.Millisecond,
			CharSet:         yacspin.CharSets[11],
			Suffix:          "",
			SuffixAutoColon: true,
			Message:         "",
			//StopCharacter:   "✔",
			StopCharacter: "🔮",
			ColorAll:      true,
			Colors:        []string{"fgCyan"},
			StopColors:    []string{"fgCyan"},
		}

		spinner, err := yacspin.New(cfg)
		if err != nil {
			panic(err)
		}

		spinner.Start()

		postbody := &PostBody{
			Model:    OPENAI_MODEL,
			Messages: messages,
		}

		body, err := json.Marshal(postbody)
		if err != nil {
			panic(err)
		}

		r, err := http.NewRequest("POST", CHAT_GPT_URL, bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", "Bearer "+apiKey)

		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			panic(err)
		}

		responseBody := &ResponseBody{}
		derr := json.NewDecoder(res.Body).Decode(responseBody)
		if derr != nil {
			panic(derr)
		}
		res.Body.Close()

		generated := responseBody.Choices[0].Message
		messages = append(messages, generated)

		spinner.Stop()
		renderMarkdown(generated.Content)
	}
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

func getStringFromFile(filename string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	file, err := os.Open(homedir + "/.config/v/chat/" + filename)
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
