package wordpress

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WPPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type WPResponse struct {
	ID int `json:"id"`
}

func PostBlog(content string, siteName string, apiKey string, postUserName string) {
	// The URL of the WordPress API endpoint
	url := siteName + "/wp-json/wp/v2/posts"

	// The data to send in the POST request
	data := map[string]interface{}{
		"title":   "Your Blog Title",
		"content": content,
		"status":  "publish",
	}

	// Convert the data to JSON
	jsonData, _ := json.Marshal(data)

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	// Add the necessary headers
	req.Header.Add("Content-Type", "application/json")

	// Basic auth
	req.SetBasicAuth(postUserName, apiKey)
	// Send the request
	client := &http.Client{}
	resp, _ := client.Do(req)

	fmt.Println("http status " + resp.Status)

	// Read and return the response
	body, _ := ioutil.ReadAll(resp.Body)
	var wpResponse WPResponse
	json.Unmarshal(body, &wpResponse)

	// Print the ID of the new post
	fmt.Printf("Posted new blog with ID: %d\n", wpResponse.ID)
}
