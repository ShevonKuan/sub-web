APP_NAME=sub-web
VUE_DIR=./frontend
GO_DIR=./

# Set environment variables
export CGO_ENABLED=0
export GIN_MODE=release

# 编译前端
.PHONY: frontend
frontend:
	cd $(VUE_DIR) && yarn install
	cd $(VUE_DIR) && yarn build

# 编译 Go 项目
.PHONY: go
go:
	cd $(GO_DIR) && go mod tidy && go build -v -o $(APP_NAME) -trimpath -ldflags "-s -w -buildid=" ./

# 打包成可执行文件
.PHONY: build
build: frontend go

# 运行
.PHONY: run
run: frontend
	cd $(GO_DIR) && go run .

# 清除编译结果
.PHONY: clean
clean:
