APP_NAME := app-service
MAIN_FILE := ./app-service/main.go
BUILD_DIR := build

GO := go
STRIP := strip

LD_FLAGS := -s -w -linkmode external -extldflags "-static"
GO_ENV := CGO_ENABLED=1 GOOS=linux GOARCH=amd64

.PHONY: release clean

release:
	@echo build static file output
	@mkdir -p $(BUILD_DIR)
	@echo "[*] Building $(APP_NAME) with full static link..."
	@$(GO_ENV) $(GO) build -o $(BUILD_DIR)/$(APP_NAME) -ldflags '$(LD_FLAGS)' $(MAIN_FILE)
	@echo "use upx compress binary file"
	@upx $(BUILD_DIR)/$(APP_NAME)
	@echo "[*] Stripping binary..."
	@echo "[✓] Build complete: $(BUILD_DIR)/$(APP_NAME)"


clean:
	@rm -rf $(BUILD_DIR)
	@echo "[*] Cleaned build artifacts"
