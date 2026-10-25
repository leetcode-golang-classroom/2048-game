.PHONY=build

build-game:
	@CGO_ENABLED=1 GOOS=linux go build -o bin/2048-game cmd/main.go

run-game: build-game
	@env -u EGL_PLATFORM ./bin/2048-game


coverage:
	@go test -v -cover ./internal/game/...

test:
	@go test -v ./internal/game/...

build-wasm:
	@env GOOS=js GOARCH=wasm go build -o web/2048.wasm ./cmd/main.go