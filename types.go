package main

type GptResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Text         string `json:"text"`
	} `json:"choices"`
}

type WPPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type WPResponse struct {
	ID int `json:"id"`
}
