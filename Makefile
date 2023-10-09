# https://www.gnu.org/software/make/manual/html_node/Quick-Reference.html
# https://stackoverflow.com/questions/448910/what-is-the-difference-between-the-gnu-makefile-variable-assignments-a

# Change these variables as necessary.
BINARY_NAME := go-playground

# 빌드 플래그
GO_FLAGS += -v
GO_FLAGS += -o=./bin/${BINARY_NAME}

# VCS 관련 meta 정보를 바이너리에 포함하지 않도록 설정.
# 컨테이너 환경에서 빌드할 때 VCS 정보를 가져오지 못해 빌드가 실패하는 경우가 있음.
# 필요하면 빌드하기 전에 git safe directory 설정을 추가하면 되긴 하는듯
# git config --global --add safe.directory ${shell pwd}
# GO_FLAGS += -buildvcs=false

# 컴파일러 플래그 설정
# 아래 옵션은 컴파일 시 최적화를 비활성화한다.
# `go tool compile`(또는 https://golang.org/cmd/compile/)를 실행하면 더 많은 옵션 정보를 얻을 수 있다.
GO_FLAGS += -gcflags="all=-N -l"

# 링커 플래그 설정
# 아래 설정은 빌드된 바이너리의 크기를 최대한 줄이기 위해 디버깅 정보를 제거하는 설정이다.
# `go tool link`(또는 https://golang.org/cmd/link/)를 실행하면 더 많은 옵션 정보를 얻을 수 있다.
# GO_FLAGS += -ldflags="-s -w"

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## audit: run quality control checks
.PHONY: audit
audit:
	@echo "GOPATH: ${shell go env GOPATH}"
	go mod tidy
	go mod verify
	go fmt ./...
	go vet ./...

## dev: run the application
.PHONY: run
run: audit
	@rm -rf ./bin
	@go build ${GO_FLAGS} ./src
	@echo "Running..."
	@./bin/${BINARY_NAME}
