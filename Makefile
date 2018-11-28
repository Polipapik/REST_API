.PHONY: build up down clean

build:
		docker build -t world_db ./db
		go build -o app/main app/main.go
		docker build -t app ./app
		rm app/main

up:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm world_db app

#remove:
