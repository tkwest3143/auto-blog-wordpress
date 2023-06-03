package openai

import (
	"auto-blog-wordpress/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GptRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type GptResponse struct {
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func GetContent(topic string, apiKey string) (string, string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	fmt.Println("---------- input settings ----------")
	file, err := os.Open("settings.json")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	var setting types.Setting
	err = decoder.Decode(&setting)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	fmt.Println(setting)
	fmt.Println("---------- end  ----------")

	fmt.Println("---------- the topics ----------")
	topicLog := fmt.Sprintf("generating blog for gpt [topic = %s]", topic)
	fmt.Println(topicLog)
	fmt.Println("---------- end  ----------")

	fmt.Println("---------- this prompt create ----------")
	conditions := strings.Join(setting.Post.Conditions, "\n")
	prompt := strings.ReplaceAll(setting.Post.Prompt, "{topic}", topic)
	prompt = strings.ReplaceAll(prompt, "{conditions}", conditions)
	fmt.Println(prompt)
	fmt.Println("---------- prompt end  ----------")
	payload := &GptRequest{
		Model: "gpt-4",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{
				Role:    "system",
				Content: fmt.Sprintf("You are a helpful assistant that speaks %s", setting.Post.Language),
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	fmt.Println("http status " + resp.Status)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var gpt4Response GptResponse
	err = json.Unmarshal(body, &gpt4Response)
	if err != nil {
		return "", "", err
	}
	text := gpt4Response.Choices[0].Message.Content
	split := strings.SplitN(text, "\n", 2)
	title := split[0]
	content := ""
	if len(split) > 1 {
		content = split[1]
	}
	fmt.Println("usage PromptTokens =", gpt4Response.Usage.PromptTokens)
	fmt.Println("usage CompletionTokens =", gpt4Response.Usage.CompletionTokens)
	fmt.Println("usage TotalTokens =", gpt4Response.Usage.TotalTokens)
	return title, content, nil
}
