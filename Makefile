BINARY_NAME := api-gym

build: $(BINARY_NAME)

$(BINARY_NAME): assets/css/tail.min.css
	go build -ldflags="-X main.buildTag=$(shell uuidgen)"

assets/css/tail.min.css: assets/css/tail.components.css
	tailwindcss -i assets/css/tail.components.css -o assets/css/tail.min.css --minify

render: build
	./api-gym render

clean:
	rm -f $(BINARY_NAME)
	rm -f assets/css/tail.min.css

.PHONY: build render clean

