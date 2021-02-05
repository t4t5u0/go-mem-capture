NAME := go-shutter

SRCS := $(shell find . -type -f -name '*.go')

bin/$(NAME): $(SRCS)
	go build  -o bin/$(NAME) main.go
