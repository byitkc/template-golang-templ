.PHONY: dev run

dev:
	@echo "Building Tailwind and DaisyUI"
	@tailwindcss -i ./static/styles/tailwind.css -o ./static/styles/dist/styles.css -v
	@echo "Building templ"
	@templ generate
	@echo "Starting development server"
	@go run cmd/main.go --logtostderr=true

run:
	@echo "Building Tailwind and DaisyUI"
	@tailwindcss -i ./static/styles/tailwind.css -o ./static/styles/dist/styles.css
	@echo "Building templ"
	@templ generate
	@echo "Starting development server"
	@go run cmd/main.go