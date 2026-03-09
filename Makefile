##
## COREL PROJECT, 2026
## Corel
## File description:
## Makefile
##

SRC = main.go
# SRC-CMD = cmd/

NAME = corel

all: $(NAME)

$(NAME): $(SRC)
	@echo "Building $(NAME)..."
	@mkdir -p bin
	@go build -o bin/$(NAME) $(SRC)

clean: bin/$(NAME)
	rm -f bin/$(NAME)
	rm -rf bin

re: clean all

.PHONY: all clean re