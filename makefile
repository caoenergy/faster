BINARY_NAME := faster

RELEASES_DIST := ./releases

.PHONY: clean
clean:
	@echo "clean releases directory..."
	@rm -rf $(RELEASES_DIST)

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ${RELEASES_DIST}/windows/amd64/${BINARY_NAME}.exe .  && upx -9 ${RELEASES_DIST}/windows/amd64/${BINARY_NAME}.exe
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ${RELEASES_DIST}/darwin/amd64/${BINARY_NAME} . && upx -9 --force-macos ${RELEASES_DIST}/darwin/amd64/${BINARY_NAME}
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ${RELEASES_DIST}/darwin/arm64/${BINARY_NAME} . && upx -9 --force-macos ${RELEASES_DIST}/darwin/arm64/${BINARY_NAME}
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ${RELEASES_DIST}/linux/amd64/${BINARY_NAME} . && upx -9 ${RELEASES_DIST}/linux/amd64/${BINARY_NAME}
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ${RELEASES_DIST}/linux/arm64/${BINARY_NAME} . && upx -9 ${RELEASES_DIST}/linux/arm64/${BINARY_NAME}
