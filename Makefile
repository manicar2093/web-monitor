run:
	@go run main.go

build_linux:
	@echo "Building to LINUX initialized :3"
	@GOOS=linux go build -o web-monitor_linux
	@echo "Build Done :D"

build_win:
	@echo "Building to WIN initialized :3"
	@GOOS=windows go build -o web-monitor_win.exe
	@echo "Build Done :D"

build_all:
	@make build_linux
	@make build_win

test_all:
	@go test ./... -v

coverage:
	@go test ./... -cover

coverage_html:
	@go test ./... -coverprofile c.out && go tool cover -html=c.out