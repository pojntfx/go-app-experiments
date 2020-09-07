build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/counter-app/main.go
	@go build -o counter-server cmd/counter-server/main.go

run: build
	@./counter-server

clean: build
	@go clean
	@-rm web/app.wasm

dev:
	@nodemon -w . -e go,css --exec "pkill -9 counter-server; make run"