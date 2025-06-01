#!/bin/sh

# run image
docker run gobel --name gobel

# copy result
docker cp gobel:/sobel_edge/new-image.jpg ./result.jpg

# remove unneeded container
docker rm gobel