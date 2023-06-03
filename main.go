package main

import (
	"auto-blog-wordpress/openai"
	"auto-blog-wordpress/wordpress"
	"bufio"
	"fmt"
	"os"
)

func main() {
	openAIKey := os.Getenv("OPENAI_API_KEY")
	siteName := os.Getenv("WORDPRESS_SITE_NAME")
	postUserName := os.Getenv("WORDPRESS_USER_NAME")
	wordpressKey := os.Getenv("WORDPRESS_API_KEY")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the topic for the blog post:")
	topic, _ := reader.ReadString('\n')

	fmt.Println("Thank you input topic")
	topicLog := fmt.Sprintf("generating blog for gpt [topic = %s\n]", topic)
	fmt.Println(topicLog)
	title, content, err := openai.GetContent(topic, openAIKey)
	fmt.Println("completing generating blog for gpt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("posting blog")
	wordpress.PostBlog(title, content, siteName, wordpressKey, postUserName)
	fmt.Println("complete posting blog")
}
