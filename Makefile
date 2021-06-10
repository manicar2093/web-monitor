run:
	@go run main.go

build_linux:
	@echo "Building to LINUX initialized :3"
	@GOOS=linux go build
	@echo "Build Done :D"

build_win:
	@echo "Building to WIN initialized :3"
	@GOOS=windows go build
	@echo "Build Done :D"

build_all:
	@make build_linux
	@make build_win