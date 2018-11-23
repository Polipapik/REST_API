.PHONY: build up down clean

build:
		docker build -t world-db ./db
		go build -o app/main app/main.go app/app.go app/model.go 
		docker build -t app ./app
		rm app/main

up:
		docker-compose up

down:
		docker-compose down

clean:
		docker rm world-db app

#remove:
