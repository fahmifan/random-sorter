BINARY		:= random-sorter

seed-number:
	@go run ./seeder/* > data.out

build:
	# @GOARCH=amd64 GOOS=windows go build -ldflags="-w -s" -o output/${BINARY}_windows.exe . && upx output/${BINARY}_windows.exe
	# @GOARCH=386 GOOS=windows go build -ldflags="-w -s" -o output/${BINARY}_windows32.exe . && upx output/${BINARY}_windows32.exe
	@GOARCH=amd64 GOOS=linux go build -ldflags="-w -s" -o output/${BINARY}_linux . && upx output/${BINARY}_linux