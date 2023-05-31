include .env
export

.PHONY: build run

build:
	go build -o build/auto-blog-wordpress

run: build
	./build/auto-blog-wordpress