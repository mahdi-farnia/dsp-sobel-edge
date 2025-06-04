# Sobel Operator On Image

Applying sobel operator on given image, using image convolution.

This project was meant to satisfy the my DSP lecture project, no security concern or
performance is considered! PR's are welcome if you want. :smiley:

## Build

You need a go compiler.
Optionally make.

Without make:

```sh
go build ./cmd/sobel_edge -o ./sobel_edge
```

with make:

```sh
make build_mac #for macintosh
make build_linux #for linux
make build_win #for windows
```

Im not a Makefile expert, so i added 3 task for different platforms.

## Copyright

Code is licensed under ISC.

Photo by [Mitul Grover](https://unsplash.com/@mitulgrover?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) on [Unsplash](https://unsplash.com/photos/orange-ferrari-car-parked-near-green-leaf-plants-during-daytime-L0MJaqt3euw?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)
