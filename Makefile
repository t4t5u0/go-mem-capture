NAME := mem-capture

SRCS := $(shell find . -type -f -name '*.go')

build:
	go build  -o bin/$(NAME) main.go && sudo cp bin/$(NAME) /bin/