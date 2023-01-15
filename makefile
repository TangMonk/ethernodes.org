BINARY_NAME=builds/ethernodes

build:
	GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o ${BINARY_NAME}-darwin .
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ${BINARY_NAME}-linux .
	GOARCH=amd64 GOOS=windows go build -ldflags="-s -w" -o ${BINARY_NAME}-windows.exe .

upx:
	upx --brute ${BINARY_NAME}-darwin
	upx --brute ${BINARY_NAME}-linux
	upx --brute ${BINARY_NAME}-windows.exe

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows