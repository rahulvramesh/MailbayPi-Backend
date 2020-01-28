build:
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build -o bin/mailbaypi .
