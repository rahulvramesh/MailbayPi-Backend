build:
	CC=gcc-arm-linux-gnueabihf CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o bin/mailbay .
