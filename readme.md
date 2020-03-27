# sockit

This is the code for the day 8 of the [Advent of Corona challenge](https://adventofcorona.hackersatupc.org/).

## Run

### A pelo

`go run .`

### Docker

- Build `docker build -t sockit .`
- Run `docker run -p 8080:8080 sockit:latest`

## Build from Docker Hub

If you are a lazy bastard, then just:

- Pull the image `docker pull tarasyarema/sockit:tagname`
- Run the container `docker run -p 8080:8080 tarasyarema/sockit:tagname`
