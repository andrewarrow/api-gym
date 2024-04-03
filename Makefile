BINARY_NAME := api-gym
DATABASE_URL := postgres://fred3:fred3@localhost/feedback?sslmode=disable

build: $(BINARY_NAME)

$(BINARY_NAME): assets/css/tail.min.css
	go build -ldflags="-X main.buildTag=$(shell uuidgen)"

assets/css/tail.min.css: assets/css/tail.components.css
	tailwindcss -i assets/css/tail.components.css -o assets/css/tail.min.css --minify

wasm: render
	cp main.go save && \
  cp wasm/main.go . && \
  GOOS=js GOARCH=wasm go build -ldflags="-s -w -X main.useLive=true" -o assets/other/json.wasm && \
	rm json.wasm.gz && \
  gzip json.wasm

render: build
	./api-gym render

run: render
	DATABASE_URL=$(DATABASE_URL) \
  echo "3" && \
  ./api-gym run 3000

clean:
	rm -f $(BINARY_NAME)
	rm -f assets/css/tail.min.css

.PHONY: build render run clean

