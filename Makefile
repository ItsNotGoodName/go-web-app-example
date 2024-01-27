export VITE_HOST=127.0.0.1
export WEB_PATH=web

-include .env

_:
	mkdir "$(WEB_PATH)/dist" -p && touch "$(WEB_PATH)/dist/keep"

generate:
	go generate ./...

run:
	go run .

preview: generate run

# Dev

dev:
	air

dev-web:
	cd "$(WEB_PATH)" && pnpm install && pnpm run dev

# Gen

gen: gen-templ

gen-templ:
	cd "$(WEB_PATH)" && templ generate

# Tooling

tooling: tooling-air tooling-templ

tooling-air:
	go install github.com/cosmtrek/air@latest

tooling-templ:
	go install github.com/a-h/templ/cmd/templ@latest
