NAME = cos
PACKAGE = github.com/fpay/cos-go-sdk-v5/cli
MAIN = $(PACKAGE)/main

BUILD_FLAGS= -mod vendor -v -o $(NAME) main/main.go

CL_RED  = "\033[0;31m"
CL_BLUE = "\033[0;34m"
CL_GREEN = "\033[0;32m"
CL_ORANGE = "\033[0;33m"
CL_NONE = "\033[0m"

define color_out
	@echo $(1)$(2)$(CL_NONE)
endef

build:
	@go mod vendor
	@go build $(BUILD_FLAGS)

linux:
	@go mod vendor
	@GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS)

.PHONY: all
all:
	build
