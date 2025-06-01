.PHONY: vet build

vet: ./cmd/sobel_edge/*.go
	go vet $<

build: vet
	go build -o ./sobel_edge ./cmd/sobel_edge
