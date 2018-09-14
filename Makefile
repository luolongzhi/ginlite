APP_NAME=ginlite

all: build 
build:
	go build -o $(APP_NAME)

clean:
	rm $(APP_NAME)

deps:
	go get -u github.com/gin-gonic/gin

