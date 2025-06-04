.PHONY: vet build

OUT = ./sobel_edge
ENTRY = ./cmd/sobel_edge

vet: ./cmd/sobel_edge/*.go
	go vet $<

build_linux: vet
	GOOS=linux GOARCH=amd64 go build -o $(OUT) $(ENTRY)

build_win: vet
	GOOS=windows GOARCH=amd64 go build -o $(OUT) $(ENTRY)

build_mac: vet
	GOOS=darwin GOARCH=arm64 go build -o $(OUT) $(ENTRY)