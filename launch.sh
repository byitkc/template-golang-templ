#!/bin/bash

if [ ! -f .env ]; then
    export $(cat .env | xargs)
fi

echo "Building tailwindcss"
tailwindcss -i ./static/styles/tailwind.css -o ./static/styles/dist/styles.css -m
echo "Building templ"
templ generate
echo "Starting development server"
go run cmd/main.go