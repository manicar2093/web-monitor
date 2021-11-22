run: clean_generates
	@go run cmd/api/main.go

clean_generates:
	@rm -r cmd/api/templates
	@rm -r cmd/api/static
	@go generate cmd/api/main.go

build_linux: clean_generates
	@echo "Building to LINUX initialized :3"
	@GOOS=linux go build -o web-monitor_linux cmd/api/main.go
	@echo "Build Done :D"

build_win: clean_generates
	@echo "Building to WIN initialized :3"
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build -o web-monitor_win.exe cmd/api/main.go
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