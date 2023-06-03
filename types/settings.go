package types

type Setting struct {
	Post Post `json:"post"`
}

type Post struct {
	Language   string   `json:"language"`
	Prompt     string   `json:"prompt"`
	Conditions []string `json:"conditions"`
}
