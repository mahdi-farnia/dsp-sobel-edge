# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /workdir
COPY . .

# building executable
RUN apk add --no-cache make
RUN make build

# Run Stage
FROM alpine

COPY --from=builder ./sobel_edge /sobel_edge/main
COPY ./assets/image.jpg /sobel_edge/image.jpg

RUN chmod +x /sobel_edge/main

CMD [ "/sobel_edge/main", "/sobel_edge/image.jpg", "/sobel_edge/new-image.jpg" ]
