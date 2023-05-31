package main

import (
	"auto-blog-wordpress/openai"
	"auto-blog-wordpress/wordpress"
	"fmt"
	"os"
)

func main() {
	openAIKey := os.Getenv("OPENAI_API_KEY")
	siteName := os.Getenv("WORDPRESS_SITE_NAME")
	postUserName := os.Getenv("WORDPRESS_USER_NAME")
	wordpressKey := os.Getenv("WORDPRESS_API_KEY")
	fmt.Println(openAIKey)
	fmt.Println(wordpressKey)

	var topic string
	fmt.Println("Enter the topic for the blog post:")
	fmt.Scanln(&topic)

	content, _ := openai.GetContent(topic, openAIKey)

	wordpress.PostBlog(content, siteName, wordpressKey, postUserName)
}
